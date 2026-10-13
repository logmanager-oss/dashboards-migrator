// Package visualization handles migration of LM3 panel visualizations to LM4 format.
package visualization

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type MigrationParams struct {
	Title             string
	Field             string
	FieldSize         int
	Queries           []lm3.Query
	Columns           []string
	Sort              []string
	VisualizationType objects.VisType
	ID                string
	Span              int
	IndexPattern      string
	Content           string
	Metric            string
	MetricField       string
	Filters           []lm4.Filter
	Colors            map[string]string
}

type LM4Visualization struct {
	SavedObject *lm4.SavedObject
	VisState    *lm4.VisState
	Search      *lm4.SearchSourceJSON
}

func MigratePanelToVisualization(params *MigrationParams) (*lm4.SavedObject, error) {
	_, isMarkdown := params.VisualizationType.(*objects.Markdown)

	vis := &LM4Visualization{
		SavedObject: params.VisualizationType.GetDefaultVisualizationSavedObject(params.IndexPattern),
		VisState:    params.VisualizationType.GetDefaultVisState(),
		Search:      objects.GetDefaultSearchObject(!isMarkdown),
	}

	// Text panels have no queries, so the "all queries" default must not apply to them.
	if isMarkdown {
		vis.VisState.Params.Markdown = params.Content
	} else if len(params.Filters) == 0 {
		vis.migrateQueries(params.Queries)
	}
	vis.migrateConfig(params)
	vis.migrateTitles(params.Title)
	vis.SavedObject.ID = params.ID

	if len(params.Filters) > 0 {
		uiState, err := json.Marshal(map[string]interface{}{"vis": map[string]interface{}{"colors": params.Colors}})
		if err != nil {
			return nil, fmt.Errorf("marshalling ui state: %w", err)
		}
		vis.SavedObject.Attributes.UIStateJSON = string(uiState)
	}

	finalVisualizationObject, err := vis.buildFinalVisualizationObject(params.VisualizationType)
	if err != nil {
		return nil, err
	}

	return finalVisualizationObject, nil
}

var (
	_rawBeforeColon = regexp.MustCompile(`\.raw(\s*:)`)
	_rawAfterExists = regexp.MustCompile(`(_exists_:\s*[^\s()]+?)\.raw\b`)
)

// StripRaw removes the LM3 ".raw" suffix from a field name.
func StripRaw(field string) string {
	return strings.TrimSuffix(field, ".raw")
}

// NormalizeLucene strips the LM3 ".raw" field suffix, which LM4 fields do not have.
func NormalizeLucene(q string) string {
	q = _rawBeforeColon.ReplaceAllString(q, "$1")
	return _rawAfterExists.ReplaceAllString(q, "$1")
}

func (vis *LM4Visualization) migrateQueries(lm3queries []lm3.Query) {
	queries := make([]string, 0, len(lm3queries))
	// Parentheses keep "a AND b" from one query from binding to the next one.
	wrap := len(lm3queries) > 1

	for _, query := range lm3queries {
		// "*" and an empty query both mean ALL events, the default search in LM4
		q := strings.TrimSpace(query.Query)
		if q == "*" || q == "" {
			return
		}

		q = NormalizeLucene(q)
		if wrap {
			q = "(" + q + ")"
		}
		queries = append(queries, q)
	}

	vis.Search.Query.Query = strings.Join(queries, " OR ")
}

func (vis *LM4Visualization) migrateConfig(params *MigrationParams) {
	cfg := objects.AggConfig{
		Field:       params.Field,
		Size:        params.FieldSize,
		Metric:      params.Metric,
		MetricField: params.MetricField,
		Filters:     params.Filters,
	}
	vis.VisState.Aggs = params.VisualizationType.GetVisualizationConfig(cfg)
	vis.SavedObject.Attributes.Columns = params.Columns
	vis.migrateSort(params.Sort)
	vis.migrateSeriesLabel(objects.MetricLabel(cfg))
}

// migrateSort keeps the saved-search sort empty for "@timestamp desc", the LM4 default,
// and for a direction OpenSearch would reject.
func (vis *LM4Visualization) migrateSort(sort []string) {
	if len(sort) != 2 {
		return
	}
	field, dir := StripRaw(sort[0]), sort[1]
	if dir != "asc" && dir != "desc" {
		return
	}
	if field == "@timestamp" && dir == "desc" {
		return
	}
	vis.SavedObject.Attributes.Sort = []interface{}{[]interface{}{field, dir}}
}

// migrateSeriesLabel sets the series and value-axis label. It does nothing without series params.
func (vis *LM4Visualization) migrateSeriesLabel(label string) {
	params := &vis.VisState.Params
	if len(params.SeriesParams) > 0 {
		if data, ok := params.SeriesParams[0]["data"].(map[string]interface{}); ok {
			data["label"] = label
		}
	}
	if len(params.ValueAxes) > 0 {
		if title, ok := params.ValueAxes[0]["title"].(map[string]interface{}); ok {
			title["text"] = label
		}
	}
}

func (vis *LM4Visualization) migrateTitles(title string) {
	vis.VisState.Title = title
	vis.SavedObject.Attributes.Title = title
}

func (vis *LM4Visualization) buildFinalVisualizationObject(visualizationType objects.VisType) (*lm4.SavedObject, error) {
	// Log Overview does not have visState object, so skip it
	if _, ok := visualizationType.(*objects.LogOverview); !ok {
		visStateRaw, err := json.Marshal(vis.VisState)
		if err != nil {
			return nil, err
		}

		vis.SavedObject.Attributes.VisState = string(visStateRaw)
	}

	searchRaw, err := json.Marshal(vis.Search)
	if err != nil {
		return nil, err
	}

	vis.SavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)

	return vis.SavedObject, nil
}
