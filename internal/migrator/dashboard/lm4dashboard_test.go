package dashboard

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

func TestLM4Dashboard_CalculateGridPosition(t *testing.T) {
	tests := []struct {
		name                string
		span                int
		currentGridPosition struct {
			X int
			Y int
		}
		expected *lm4.GridData
	}{
		{
			name: "Test calculate grid position - Test case 1",
			span: 12,
			// Current grid position is at 0:0 meaning no other visualization exists yet
			currentGridPosition: struct {
				X int
				Y int
			}{
				X: 0,
				Y: 0,
			},
			expected: &lm4.GridData{
				X: 0,
				Y: 0,
				W: 48,
				H: 15,
			},
		},
		{
			name: "Test calculate grid position - Test case 2",
			span: 12,
			// Current grid position is at 48:15 meaning there already is a visualization which takes up entire row
			currentGridPosition: struct {
				X int
				Y int
			}{
				X: 48,
				Y: 15,
			},
			// As such we expect grid calculation to return coordinates below existing visualization
			expected: &lm4.GridData{
				X: 0,
				Y: 30,
				W: 48,
				H: 15,
			},
		},
		{
			name: "Test calculate grid position - Test case 3",
			span: 4,
			// Current grid position is at 16:15 meaning there is already a visualization which takes up 1/3 of row
			currentGridPosition: struct {
				X int
				Y int
			}{
				X: 16,
				Y: 15,
			},
			// As such we expect grid calculation to return coordinates next to existing visualization
			expected: &lm4.GridData{
				X: 16,
				Y: 15,
				W: 16,
				H: 15,
			},
		},
		{
			name: "Test calculate grid position - Test case 4",
			span: 12,
			// Current grid position is at 16:15 meaning there is already a visualization which takes up 1/3 of row
			currentGridPosition: struct {
				X int
				Y int
			}{
				X: 16,
				Y: 15,
			},
			// Because next visualization has a span of 12 it's not going to fit next to existing one
			// As such we expect grid calculation to return coordinates below existing visualization
			expected: &lm4.GridData{
				X: 0,
				Y: 30,
				W: 48,
				H: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dashboard := &LM4Dashboard{
				savedObject:         defaults.GetDefaultDashboardSavedObject(),
				search:              defaults.GetDefaultSearch(false),
				id:                  uuid.New().String(),
				currentGridPosition: tt.currentGridPosition,
			}
			grid := dashboard.CalculateGridPosition(tt.span)

			assert.Equal(t, tt.expected, grid)
		})
	}
}

func TestLM4Dashboard_BuildPanelObject(t *testing.T) {
	tests := []struct {
		name                string
		span                int
		id                  string
		title               string
		currentGridPosition struct {
			X int
			Y int
		}
		expected *lm4.PanelJSON
	}{
		{

			name:  "Test build panel object - test case 1",
			span:  12,
			id:    "test",
			title: "test",
			currentGridPosition: struct {
				X int
				Y int
			}{
				X: 0,
				Y: 0,
			},
			expected: &lm4.PanelJSON{
				Version:    "2.19.1",
				GridData:   lm4.GridData{X: 0, Y: 0, W: 48, H: 15, I: "test"},
				PanelIndex: "test",
				EmbeddableConfig: lm4.EmbeddableConfig{
					HidePanelTitles: false,
				},
				Title:        "test",
				PanelRefName: "panel_0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dashboard := &LM4Dashboard{
				savedObject:         defaults.GetDefaultDashboardSavedObject(),
				search:              defaults.GetDefaultSearch(false),
				id:                  uuid.New().String(),
				currentGridPosition: tt.currentGridPosition,
			}
			grid := dashboard.CalculateGridPosition(tt.span)

			panel := dashboard.BuildPanelObject(
				grid,
				tt.id,
				tt.title,
			)

			assert.Equal(t, tt.expected, panel)
		})
	}
}

func TestLM4Dashboard_BuildReferenceObject(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		refType  string
		expected *lm4.Reference
	}{
		{

			name:    "Test build panel object - test case 1",
			id:      "test",
			refType: "visualization",
			expected: &lm4.Reference{
				ID:   "test",
				Name: "panel_0",
				Type: "visualization",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dashboard := &LM4Dashboard{
				savedObject: defaults.GetDefaultDashboardSavedObject(),
				search:      defaults.GetDefaultSearch(false),
				id:          uuid.New().String(),
			}

			panel := dashboard.BuildReferenceObject(
				tt.id,
				tt.refType,
			)

			assert.Equal(t, tt.expected, panel)
		})
	}
}
