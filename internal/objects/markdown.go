package objects

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type Markdown struct{}

func (m *Markdown) GetDefaultVisualizationSavedObject(string) *lm4.SavedObject {
	return &lm4.SavedObject{
		Attributes: lm4.Attributes{
			UIStateJSON: "{}",
			Version:     1,
		},
		MigrationVersion: map[string]any{"visualization": "7.10.0"},
		References:       []lm4.Reference{},
		Type:             "visualization",
		UpdatedAt:        time.Time{},
	}
}

func (m *Markdown) GetVisualizationConfig(string, int) []lm4.VisStateAggs {
	return []lm4.VisStateAggs{}
}

func (m *Markdown) GetDefaultVisState() *lm4.VisState {
	return &lm4.VisState{
		Type: "markdown",
		Aggs: []lm4.VisStateAggs{},
		Params: lm4.VisStateParams{
			Times:    []any{},
			FontSize: 12,
		},
	}
}
