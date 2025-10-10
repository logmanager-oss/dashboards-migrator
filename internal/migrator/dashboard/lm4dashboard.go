package dashboard

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
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
	currentPanelNumber int
}

func NewLM4Dashboard() *LM4Dashboard {
	return &LM4Dashboard{
		savedObject: objects.GetDefaultDashboardSavedObject(),
		search:      objects.GetDefaultSearchObject(false),
		id:          uuid.New().String(),
	}
}

func (dashboard *LM4Dashboard) BuildDashboardObject() (*lm4.SavedObject, error) {
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

func (dashboard *LM4Dashboard) LinkVisualizationToDashboardObject(params *visualization.MigrationParams, visualizationType string) {
	grid := dashboard.calculateGridPosition(params.Span)

	panel := dashboard.buildPanelObject(
		grid,
		params.ID,
		params.Title,
	)

	dashboard.Panels = append(dashboard.Panels, *panel)

	ref := dashboard.buildReferenceObject(
		params.ID,
		visualizationType,
	)

	dashboard.References = append(dashboard.References, *ref)

	dashboard.currentPanelNumber++
}

func (dashboard *LM4Dashboard) calculateGridPosition(span int) *lm4.GridData {
	width := span * 4

	if dashboard.currentGridPosition.X == objects.MaxRowWidth || dashboard.currentGridPosition.X+width > objects.MaxRowWidth {
		dashboard.currentGridPosition.X = 0
		dashboard.currentGridPosition.Y += objects.DefaultVisHeight
	}

	gridData := &lm4.GridData{
		X: dashboard.currentGridPosition.X,
		Y: dashboard.currentGridPosition.Y,
	}

	gridData.W = width
	gridData.H = objects.DefaultVisHeight
	dashboard.currentGridPosition.X += gridData.W

	return gridData
}

func (dashboard *LM4Dashboard) buildPanelObject(grid *lm4.GridData, id string, title string) *lm4.PanelJSON {
	panel := objects.GetDefaultPanelObject()
	panel.GridData = *grid
	panel.GridData.I = id
	panel.PanelIndex = id
	panel.Title = title
	panel.PanelRefName = fmt.Sprintf("panel_%d", dashboard.currentPanelNumber)

	return panel
}

func (dashboard *LM4Dashboard) buildReferenceObject(id string, refType string) *lm4.Reference {
	ref := objects.GetDefaultReferenceObject()
	ref.Name = fmt.Sprintf("panel_%d", dashboard.currentPanelNumber)
	ref.ID = id
	ref.Type = refType

	return ref
}
