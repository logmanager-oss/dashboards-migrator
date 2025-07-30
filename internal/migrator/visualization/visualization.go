package visualization

import (
	"encoding/json"
	"slices"

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
}

func NewLM4Visualisation(title string, lm3filters []lm3.Query, lm3filterIDs []int, visType vistypes.VisType) (*lm4.SavedObject, error) {
	vis := &LM4Visualization{
		SavedObject:       defaults.GetDefaultVisualizationSavedObject(),
		VisState:          defaults.GetDefaultHistogramVisState(),
		Search:            defaults.GetDefaultSearch(true),
		VisualizationType: visType,
		ID:                uuid.New().String(),
		Title:             title,
	}

	vis.setFilters(lm3filters, lm3filterIDs)
	vis.setTitles()
	vis.setVisualizationConfig(vis.VisualizationType.GetVisualizationConfig(vis.filters))

	finalVisualizationObject, err := vis.buildFinalVisualizationObject()
	if err != nil {
		return nil, err
	}

	return finalVisualizationObject, nil
}

func (vis *LM4Visualization) setFilters(lm3filters []lm3.Query, filterIDs []int) {
	for _, lm3filter := range lm3filters {
		if slices.Contains(filterIDs, lm3filter.ID) {
			vis.filters = append(vis.filters, lm4.Filter{
				Input: map[string]string{"query": lm3filter.Query, "language": "kuery"},
				Label: "",
			})
		}
	}
}

func (vis *LM4Visualization) setVisualizationConfig(visStateAggs []lm4.VisStateAggs) {
	vis.VisState.Aggs = visStateAggs
}

func (vis *LM4Visualization) setTitles() {
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
