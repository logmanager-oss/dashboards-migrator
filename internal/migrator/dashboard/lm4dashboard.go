package dashboard

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type LM4Dashboard struct {
	savedObject         *lm4.SavedObject
	search              *lm4.SearchSourceJSON
	Panels              []lm4.PanelJSON
	currentRowHeight    int
	currentGridPosition struct {
		X int
		Y int
	}
	references       []lm4.Reference
	currentRefNumber int
}

func NewLM4Dashboard() *LM4Dashboard {
	return &LM4Dashboard{
		savedObject:      objects.GetDefaultDashboardSavedObject(),
		search:           objects.GetDefaultSearchObject(false),
		Panels:           []lm4.PanelJSON{},
		references:       []lm4.Reference{},
		currentRowHeight: objects.DefaultVisHeight,
	}
}

func (dashboard *LM4Dashboard) MigrateDashboard(title string, filters []lm3.GlobalFilter, indexPattern string) (*lm4.SavedObject, error) {
	gridRaw, err := json.Marshal(dashboard.Panels)
	if err != nil {
		return nil, err
	}

	dashboard.search.Filter = dashboard.migrateFilters(filters, indexPattern)
	searchRaw, err := json.Marshal(dashboard.search)
	if err != nil {
		return nil, err
	}

	dashboard.savedObject.Attributes.Title = title
	dashboard.savedObject.Attributes.PanelsJSON = string(gridRaw)
	dashboard.savedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)
	dashboard.savedObject.References = dashboard.references

	return dashboard.savedObject, nil
}

func (dashboard *LM4Dashboard) LinkVisualizationToDashboardObject(params *visualization.MigrationParams, visualizationType string) {
	grid := dashboard.calculateGridPosition(params.Span)

	panel := dashboard.buildPanelObject(
		grid,
		params.ID,
		params.Title,
	)

	dashboard.Panels = append(dashboard.Panels, *panel)

	ref := dashboard.buildVisualizationReferenceObject(
		params.ID,
		visualizationType,
	)

	dashboard.references = append(dashboard.references, *ref)

	dashboard.currentRefNumber++
}

// StartRow makes the next visualization start on a new line, as every LM3 row does, and sets its height.
func (dashboard *LM4Dashboard) StartRow(height string) {
	if dashboard.currentGridPosition.X > 0 {
		dashboard.newLine()
	}
	dashboard.currentRowHeight = rowHeight(height)
}

func (dashboard *LM4Dashboard) newLine() {
	dashboard.currentGridPosition.X = 0
	dashboard.currentGridPosition.Y += dashboard.currentRowHeight
}

// rowHeight converts an LM3 row height such as "150px" to OSD grid rows.
func rowHeight(px string) int {
	n, err := strconv.Atoi(strings.TrimSuffix(px, "px"))
	if err != nil {
		return objects.DefaultVisHeight
	}

	return max((n+objects.GridRowPx-1)/objects.GridRowPx, objects.MinVisHeight)
}

func (dashboard *LM4Dashboard) calculateGridPosition(span int) *lm4.GridData {
	width := span * 4

	if dashboard.currentGridPosition.X == objects.MaxRowWidth || dashboard.currentGridPosition.X+width > objects.MaxRowWidth {
		dashboard.newLine()
	}

	gridData := &lm4.GridData{
		X: dashboard.currentGridPosition.X,
		Y: dashboard.currentGridPosition.Y,
	}

	gridData.W = width
	gridData.H = dashboard.currentRowHeight
	dashboard.currentGridPosition.X += gridData.W

	return gridData
}

func (dashboard *LM4Dashboard) buildPanelObject(grid *lm4.GridData, id string, title string) *lm4.PanelJSON {
	panel := objects.GetDefaultPanelObject()
	panel.GridData = *grid
	panel.GridData.I = id
	panel.PanelIndex = id
	panel.Title = title
	panel.PanelRefName = fmt.Sprintf("panel_%d", dashboard.currentRefNumber)

	return panel
}

func (dashboard *LM4Dashboard) buildVisualizationReferenceObject(id string, refType string) *lm4.Reference {
	ref := objects.GetDefaultReferenceObject()
	ref.Name = fmt.Sprintf("panel_%d", dashboard.currentRefNumber)
	ref.ID = id
	ref.Type = refType

	return ref
}

func (dashboard *LM4Dashboard) migrateFilters(lm3filters []lm3.GlobalFilter, indexPattern string) []lm4.GlobalFilter {
	lm4filters := make([]lm4.GlobalFilter, 0, len(lm3filters))
	var either []lm3.GlobalFilter

	for _, lm3filter := range lm3filters {
		if lm3filter.Mandate == "either" {
			either = append(either, lm3filter)
			continue
		}
		lm4filters = dashboard.appendFilter(lm4filters, buildFilter(lm3filter), indexPattern)
	}

	return dashboard.appendFilter(lm4filters, buildEitherFilter(either), indexPattern)
}

// appendFilter numbers the index reference by the filter's position, which is what OSD expects.
func (dashboard *LM4Dashboard) appendFilter(filters []lm4.GlobalFilter, filter *lm4.GlobalFilter, indexPattern string) []lm4.GlobalFilter {
	if filter == nil {
		return filters
	}

	indexRefName := fmt.Sprintf("kibanaSavedObjectMeta.searchSourceJSON.filter[%d].meta.index", len(filters))
	filter.Meta.IndexRefName = indexRefName
	dashboard.references = append(dashboard.references, *dashboard.buildFilterReferenceObject(indexPattern, indexRefName))

	return append(filters, *filter)
}

func buildFilter(f lm3.GlobalFilter) *lm4.GlobalFilter {
	filter := objects.GetDefaultFilterObject("")
	field := visualization.StripRaw(f.Field)

	switch f.Type {
	case "querystring":
		setQueryString(filter, visualization.NormalizeLucene(f.Query))
	case "exists":
		filter.Exists = &lm4.ExistsFilter{Field: field}
		filter.Meta.Type = "exists"
		filter.Meta.Key = field
		filter.Meta.Value = "exists"
	case "field":
		fieldFilter(filter, f, field)
	case "terms":
		setPhrase(filter, field, string(f.Value))
	default:
		slog.Warn("skipping global filter of unknown type", "type", f.Type, "id", f.ID)
		return nil
	}

	if f.Alias != "" {
		filter.Meta.Alias = f.Alias
	}
	filter.Meta.Negate = f.Mandate == "mustNot"
	filter.Meta.Disabled = !f.Active

	return filter
}

func fieldFilter(filter *lm4.GlobalFilter, f lm3.GlobalFilter, field string) {
	if isQuotedPhrase(f.Query) {
		setPhrase(filter, field, strings.Trim(f.Query, `"`))
		return
	}
	if strings.ContainsAny(f.Query, "\"*? \t()") {
		setQueryString(filter, luceneClause(f))
		return
	}
	setPhrase(filter, field, f.Query)
}

func isQuotedPhrase(q string) bool {
	return len(q) >= 2 && q[0] == '"' && q[len(q)-1] == '"' && strings.Count(q, `"`) == 2
}

func buildEitherFilter(either []lm3.GlobalFilter) *lm4.GlobalFilter {
	var clauses, keys []string
	for _, f := range either {
		if !f.Active {
			continue
		}
		clause := luceneClause(f)
		if clause == "" {
			slog.Warn("skipping global filter of unknown type", "type", f.Type, "id", f.ID)
			continue
		}
		clauses = append(clauses, clause)
		if f.Field != "" {
			keys = append(keys, visualization.StripRaw(f.Field))
		}
	}
	if len(clauses) == 0 {
		return nil
	}

	filter := objects.GetDefaultFilterObject("")
	setQueryString(filter, strings.Join(clauses, " OR "))
	filter.Meta.Alias = "either: " + strings.Join(keys, ", ")

	return filter
}

func luceneClause(f lm3.GlobalFilter) string {
	field := visualization.StripRaw(f.Field)
	switch f.Type {
	case "terms":
		return fmt.Sprintf("%s:%q", field, string(f.Value))
	case "exists":
		return "_exists_:" + field
	case "querystring":
		return "(" + visualization.NormalizeLucene(f.Query) + ")"
	case "field":
		return fmt.Sprintf("%s:(%s)", field, visualization.NormalizeLucene(f.Query))
	default:
		return ""
	}
}

func setPhrase(filter *lm4.GlobalFilter, field, value string) {
	filter.Query = &lm4.GlobalFilterQuery{MatchPhrase: map[string]string{field: value}}
	filter.Meta.Params = &lm4.FilterParams{Query: value}
	filter.Meta.Type = "phrase"
	filter.Meta.Key = field
}

func setQueryString(filter *lm4.GlobalFilter, query string) {
	filter.Query = &lm4.GlobalFilterQuery{QueryString: map[string]string{"query": query}}
	filter.Meta.Type = "query_string"
	filter.Meta.Key = "query"
}

func (dashboard *LM4Dashboard) buildFilterReferenceObject(id string, indexRefName string) *lm4.Reference {
	ref := objects.GetDefaultReferenceObject()
	ref.Name = indexRefName
	ref.ID = id
	ref.Type = "index-pattern"

	return ref
}
