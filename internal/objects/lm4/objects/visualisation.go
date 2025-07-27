package objects

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/defaults"
	vistypes "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/visTypes"
)

type VisualizationObject struct {
	SavedObject *lm4.SavedObject
	VisState    *lm4.VisState
	Search      *lm4.SearchSourceJSON
	Type        vistypes.VisType
	ID          string
	Title       string
	RefType     string
}

func NewVisualisation(visType vistypes.VisType) *VisualizationObject {
	return &VisualizationObject{
		SavedObject: defaults.GetDefaultVisualizationSavedObject(),
		VisState:    defaults.GetDefaultHistogramVisState(),
		Type:        visType,
		ID:          uuid.New().String(),
		RefType:     "visualization",
	}
}

func (v *VisualizationObject) SetVisStateAggs(visStateAggs []lm4.VisStateAggs) {
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
