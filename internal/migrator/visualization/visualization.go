package visualization

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type LM4Visualization struct {
	SavedObject       *lm4.SavedObject
	VisState          *lm4.VisState
	Search            *lm4.SearchSourceJSON
	VisualizationType vistypes.VisType
	ID                string
	Title             string
	filters           []lm4.Filter
	field             string
	size              int
}

func NewLM4Visualisation(title string, lm3filters []lm3.Filter, visType vistypes.VisType) (*lm4.SavedObject, error) {
	vis := &LM4Visualization{
		SavedObject:       defaults.GetDefaultVisualizationSavedObject(),
		VisState:          defaults.GetDefaultHistogramVisState(),
		Search:            defaults.GetDefaultSearch(true),
		VisualizationType: visType,
		ID:                uuid.New().String(),
		Title:             title,
	}

	vis.migrateFilters(lm3filters)
	vis.migrateTitles()
	vis.migrateVisualizationConfig(vis.VisualizationType.GetVisualizationConfig(vis.filters, vis.field, vis.size))

	finalVisualizationObject, err := vis.buildFinalVisualizationObject()
	if err != nil {
		return nil, err
	}

	return finalVisualizationObject, nil
}

func (vis *LM4Visualization) migrateFilters(filters []lm3.Filter) {
	for _, filter := range filters {
		// is visualisation type is EventsOverTimeWithFilters but one of the filters is topN then skip it - we do not support such visualisation yet
		if _, ok := vis.VisualizationType.(*vistypes.EventsOverTimeWithFilters); ok {
			if filter.Type == "topN" {
				continue
			}
		}
		// is visualisation type is EventsOverTimeAsSplitSeries then we must grab it's field and size to be able to create visualisation config
		if _, ok := vis.VisualizationType.(*vistypes.EventsOverTimeAsSplitSeries); ok {
			if filter.Type == "topN" {
				vis.field = filter.Field
				vis.size = filter.Size
			}
		}

		vis.filters = append(vis.filters, lm4.Filter{
			Input: map[string]string{"query": filter.Query, "language": "kuery"},
			Label: "",
		})
	}
}

func (vis *LM4Visualization) migrateVisualizationConfig(visStateAggs []lm4.VisStateAggs) {
	vis.VisState.Aggs = visStateAggs
}

func (vis *LM4Visualization) migrateTitles() {
	vis.SavedObject.Attributes.Title = vis.Title
	vis.VisState.Title = vis.Title
}

func (vis *LM4Visualization) buildFinalVisualizationObject() (*lm4.SavedObject, error) {
	visStateRaw, err := json.Marshal(vis.VisState)
	if err != nil {
		return nil, err
	}

	searchRaw, err := json.Marshal(vis.Search)
	if err != nil {
		return nil, err
	}

	vis.SavedObject.Attributes.VisState = string(visStateRaw)
	vis.SavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)

	return vis.SavedObject, nil
}
