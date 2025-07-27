package lm4objects

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/defaults"
)

type VisualizationObject struct {
	SavedObject *lm4.SavedObject
	VisState    *lm4.VisState
	Search      *lm4.SearchSourceJSON
	ID          string
	Title       string
}

func NewVisualisation() *VisualizationObject {
	return &VisualizationObject{
		SavedObject: defaults.GetDefaultVisualizationSavedObject(),
		VisState:    defaults.GetDefaultHistogramVisState(),
		ID:          uuid.New().String(),
	}
}

func (v *VisualizationObject) SetVisualizationConfig(visStateAggs []lm4.VisStateAggs) {
	v.VisState.Aggs = visStateAggs
}

func (v *VisualizationObject) SetSearch(search *lm4.SearchSourceJSON) {
	v.Search = search
}

func (v *VisualizationObject) SetTitle(title string) {
	v.Title = title
	v.SavedObject.Attributes.Title = title
	v.VisState.Title = title
}

func (v *VisualizationObject) GetFinalVisualizationObject() (*lm4.SavedObject, error) {
	visStateRaw, err := json.Marshal(v.VisState)
	if err != nil {
		return nil, err
	}

	searchRaw, err := json.Marshal(v.Search)
	if err != nil {
		return nil, err
	}

	v.SavedObject.Attributes.VisState = string(visStateRaw)
	v.SavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)

	return v.SavedObject, nil
}
