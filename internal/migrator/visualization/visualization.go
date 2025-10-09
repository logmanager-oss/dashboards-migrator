package visualization

import (
	"encoding/json"
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
	VisualizationType objects.VisType
	ID                string
	Span              int
}

type LM4Visualization struct {
	SavedObject *lm4.SavedObject
	VisState    *lm4.VisState
	Search      *lm4.SearchSourceJSON
}

func MigratePanelToVisualization(params *MigrationParams) (*lm4.SavedObject, error) {
	vis := &LM4Visualization{
		SavedObject: params.VisualizationType.GetDefaultVisualizationSavedObject(),
		VisState:    params.VisualizationType.GetDefaultVisState(),
		Search:      objects.GetDefaultSearchObject(true),
	}

	vis.migrateQueries(params.Queries)
	vis.migrateConfig(params.Field, params.FieldSize, params.Columns, params.VisualizationType)
	vis.migrateTitles(params.Title)
	vis.SavedObject.ID = params.ID

	finalVisualizationObject, err := vis.buildFinalVisualizationObject(params.VisualizationType)
	if err != nil {
		return nil, err
	}

	return finalVisualizationObject, nil
}

func (vis *LM4Visualization) migrateQueries(lm3queries []lm3.Query) {
	var queries []string

	for _, query := range lm3queries {
		// if one of the queries is * then just return as there is no reason to migrate - "*" means ALL events, which is a default search in LM4
		if query.Query == "*" {
			return
		}

		// if one of the queries is topN then skip it - we either don't support it or it's Events Over Time As Split Series so there is no need to migrate query
		if query.Type == "topN" {
			continue
		}

		queries = append(queries, query.Query)
	}

	vis.Search.Query.Query = strings.Join(queries, " or ")
}

func (vis *LM4Visualization) migrateConfig(field string, size int, columns []string, visualizationType objects.VisType) {
	vis.VisState.Aggs = visualizationType.GetVisualizationConfig(field, size)
	vis.SavedObject.Attributes.Columns = columns
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
