package migrator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func Test_gatherMigrationParams(t *testing.T) {
	tests := []struct {
		name        string
		panel       *lm3.Panel
		queries     []lm3.Query
		expected    *visualization.MigrationParams
		expectedErr error
	}{
		{
			name: "Test params gathering - test case 1 - event over time",
			panel: &lm3.Panel{
				Span: 12,
				Type: "histogram",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Title: "Events over time",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title: "Events over time",
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				VisualizationType: &objects.EventsOverTime{},
				Span:              12,
			},
		},
		{
			name: "Test params gathering - test case 2 - event over time as split series",
			panel: &lm3.Panel{
				Span: 12,
				Type: "histogram",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Title: "Events over time",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "topN",
					Query: "*",
					Field: "msg.protocol",
					Size:  5,
				},
			},
			expected: &visualization.MigrationParams{
				Title:     "Events over time",
				Field:     "msg.protocol",
				FieldSize: 5,
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "topN",
						Query: "*",
						Field: "msg.protocol",
						Size:  5,
					},
				},
				VisualizationType: &objects.EventsOverTimeAsSplitSeries{},
				Span:              12,
			},
		},
		{
			name: "Test params gathering - test case 3 - log overview",
			panel: &lm3.Panel{
				Span: 12,
				Type: "table",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Title: "All events",
				Fields: []string{
					"@timestamp",
					"meta.src.ip",
					"raw",
				},
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title: "All events",
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				Columns: []string{
					"@timestamp",
					"meta.src.ip",
					"raw",
				},
				VisualizationType: &objects.LogOverview{},
				Span:              12,
			},
		},
		{
			name: "Test params gathering - test case 4 - map",
			panel: &lm3.Panel{
				Span: 6,
				Type: "map",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Field: "msg.src_ip@ip.country_code",
				Size:  100,
				Title: "World map",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title:     "World map",
				Field:     "msg.src_ip@ip.country_code",
				FieldSize: 100,
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				VisualizationType: &objects.Map{},
				Span:              6,
			},
		},
		{
			name: "Test params gathering - test case 5 - vertical graph",
			panel: &lm3.Panel{
				Span:  4,
				Type:  "terms",
				Chart: "bar",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Field: "msg.dst_ip@ip.value",
				Size:  10,
				Title: "Destination IP",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title:     "Destination IP",
				Field:     "msg.dst_ip@ip.value",
				FieldSize: 10,
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				VisualizationType: &objects.VerticalGraph{},
				Span:              4,
			},
		},
		{
			name: "Test params gathering - test case 6 - pie graph",
			panel: &lm3.Panel{
				Span:  4,
				Type:  "terms",
				Chart: "pie",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Field: "msg.dst_ip@ip.value",
				Size:  10,
				Title: "Destination IP",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title:     "Destination IP",
				Field:     "msg.dst_ip@ip.value",
				FieldSize: 10,
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				VisualizationType: &objects.PieGraph{},
				Span:              4,
			},
		},
		{
			name: "Test params gathering - test case 7 - table graph",
			panel: &lm3.Panel{
				Span:  4,
				Type:  "terms",
				Chart: "table",
				Queries: lm3.Queries{
					Mode: "all",
					IDs:  []int{0},
				},
				Field: "msg.dst_ip@ip.value",
				Size:  10,
				Title: "Destination IP",
			},
			queries: []lm3.Query{
				{
					ID:    0,
					Type:  "lucene",
					Query: "*",
				},
			},
			expected: &visualization.MigrationParams{
				Title:     "Destination IP",
				Field:     "msg.dst_ip@ip.value",
				FieldSize: 10,
				Queries: []lm3.Query{
					{
						ID:    0,
						Type:  "lucene",
						Query: "*",
					},
				},
				VisualizationType: &objects.TableGraph{},
				Span:              4,
			},
		},
		{
			name: "Test params gathering - test case 8 - unknown panel type",
			panel: &lm3.Panel{
				Type: "unknown",
			},
			expectedErr: &PanelTypeNotFoundError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := gatherMigrationParams(tt.panel, tt.queries)
			if err != nil {
				if assert.ErrorAs(t, err, &tt.expectedErr) {
					return
				}
				t.Fatal(err)
			}

			// Inject ID since it generated inside gather function
			tt.expected.ID = params.ID

			assert.Equal(t, tt.expected, params)
		})
	}
}
