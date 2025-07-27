package lm4objects

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/defaults"
)

type DashboardObject struct {
	savedObject           *lm4.SavedObject
	search                *lm4.SearchSourceJSON
	references            []lm4.Reference
	grid                  []lm4.PanelJSON
	id                    string
	currentVisCoordinates struct {
		X int
		Y int
	}
	currentPanelNumber int
}

func NewDashboard() *DashboardObject {
	return &DashboardObject{
		savedObject: defaults.GetDefaultVisualizationSavedObject(),
		id:          uuid.New().String(),
	}
}

func (d *DashboardObject) SetSearch(search *lm4.SearchSourceJSON) {
	d.search = search
}

func (d *DashboardObject) SetAndAppendGridData(w int, h int, id string, title string) {
	grid := defaults.GetDefaultGridData()
	grid.GridData = d.calculateCurrentCoordinates(w, h)
	grid.GridData.I = id
	grid.PanelIndex = id
	grid.Title = title
	grid.PanelRefName = fmt.Sprintf("panel_%d", d.currentPanelNumber)

	d.grid = append(d.grid, *grid)
}

func (d *DashboardObject) SetAndAppendReference(id string, refType string) {
	ref := defaults.GetDefaultReference()
	ref.ID = id
	ref.Name = fmt.Sprintf("panel_%d", d.currentPanelNumber)
	ref.Type = refType
	d.references = append(d.references, ref)

	d.currentPanelNumber++
}

func (d *DashboardObject) calculateCurrentCoordinates(w int, h int) lm4.GridData {
	gridData := lm4.GridData{
		X: d.currentVisCoordinates.X,
		Y: d.currentVisCoordinates.Y,
	}

	if d.currentVisCoordinates.X == 48 {
		d.currentVisCoordinates.X = 0
		d.currentVisCoordinates.Y += 15
	}

	gridData.W = w
	gridData.H = h
	d.currentVisCoordinates.X += gridData.W

	return gridData
}

func (d *DashboardObject) GetFinalDashboardObject() (*lm4.SavedObject, error) {
	gridRaw, err := json.Marshal(d.grid)
	if err != nil {
		return nil, err
	}

	searchRaw, err := json.Marshal(d.search)
	if err != nil {
		return nil, err
	}

	d.savedObject.Attributes.PanelsJSON = string(gridRaw)
	d.savedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)

	return d.savedObject, nil
}
