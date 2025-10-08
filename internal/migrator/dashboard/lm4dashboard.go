package dashboard

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type LM4Dashboard struct {
	savedObject         *lm4.SavedObject
	search              *lm4.SearchSourceJSON
	Panels              []lm4.PanelJSON
	id                  string
	currentGridPosition struct {
		X int
		Y int
	}
	References         []lm4.Reference
	CurrentPanelNumber int
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
	dashboard.savedObject.References = dashboard.References

	return dashboard.savedObject, nil
}

func (dashboard *LM4Dashboard) BuildPanelObject(grid *lm4.GridData, id string, title string) *lm4.PanelJSON {
	panel := defaults.GetDefaultPanel()
	panel.GridData = *grid
	panel.GridData.I = id
	panel.PanelIndex = id
	panel.Title = title
	panel.PanelRefName = fmt.Sprintf("panel_%d", dashboard.CurrentPanelNumber)

	return panel
}

func (dashboard *LM4Dashboard) CalculateGridPosition(span int) *lm4.GridData {
	width := span * 4

	if dashboard.currentGridPosition.X == defaults.MaxRowWidth || dashboard.currentGridPosition.X+width > defaults.MaxRowWidth {
		dashboard.currentGridPosition.X = 0
		dashboard.currentGridPosition.Y += defaults.DefaultVisHeight
	}

	gridData := &lm4.GridData{
		X: dashboard.currentGridPosition.X,
		Y: dashboard.currentGridPosition.Y,
	}

	gridData.W = width
	gridData.H = defaults.DefaultVisHeight
	dashboard.currentGridPosition.X += gridData.W

	return gridData
}

func (dashboard *LM4Dashboard) BuildReferenceObject(id string, refType string) *lm4.Reference {
	ref := defaults.GetDefaultReference()
	ref.Name = fmt.Sprintf("panel_%d", dashboard.CurrentPanelNumber)
	ref.ID = id
	ref.Type = refType

	return ref
}
