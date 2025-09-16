package dashboard

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type LM4Dashboard struct {
	savedObject *lm4.SavedObject
	search      *lm4.SearchSourceJSON
	references  []lm4.Reference
	Panels      []lm4.PanelJSON
	id          string
}

func NewLM4Dashboard() *LM4Dashboard {
	return &LM4Dashboard{
		savedObject: defaults.GetDefaultDashboardSavedObject(),
		search:      defaults.GetDefaultSearch(false),
		id:          uuid.New().String(),
	}
}

func (dashboard *LM4Dashboard) BuildFinalDashboardObject() (*lm4.SavedObject, error) {
	gridRaw, err := json.Marshal(dashboard.Panels)
	if err != nil {
		return nil, err
	}

	searchRaw, err := json.Marshal(dashboard.search)
	if err != nil {
		return nil, err
	}

	dashboard.savedObject.Attributes.PanelsJSON = string(gridRaw)
	dashboard.savedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)
	dashboard.savedObject.References = dashboard.references

	return dashboard.savedObject, nil
}
