package migrator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func Test_gatherMigrationParams(t *testing.T) {
	tests := []struct {
		name        string
		panel       *lm3.Panel
		queries     []lm3.Query
		expected    *visualization.MigrationParams
		expectedErr *PanelTypeNotFoundError
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
				Metric:            "count",
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
				Metric:            "count",
				Span:              12,
			},
		},
		{
			name: "Test params gathering - event over time with several queries",
			panel: &lm3.Panel{
				Span:    12,
				Type:    "histogram",
				Queries: lm3.Queries{Mode: "all", IDs: []int{0, 1}},
				Title:   "Events over time",
			},
			queries: []lm3.Query{
				{ID: 0, Type: "lucene", Query: "msg.status:failed", Alias: "failed", Color: "#aaaaaa"},
				{ID: 1, Type: "lucene", Query: "meta.parser.raw:lm-fortigate", Color: "#bbbbbb"},
				{ID: 2, Type: "lucene", Query: "", Color: "#cccccc"},
			},
			expected: &visualization.MigrationParams{
				Title: "Events over time",
				Queries: []lm3.Query{
					{ID: 0, Type: "lucene", Query: "msg.status:failed", Alias: "failed", Color: "#aaaaaa"},
					{ID: 1, Type: "lucene", Query: "meta.parser.raw:lm-fortigate", Color: "#bbbbbb"},
					{ID: 2, Type: "lucene", Query: "", Color: "#cccccc"},
				},
				VisualizationType: &objects.EventsOverTime{},
				Metric:            "count",
				Span:              12,
				Filters: []lm4.Filter{
					{Input: map[string]string{"query": "msg.status:failed", "language": "lucene"}, Label: "failed"},
					{Input: map[string]string{"query": "meta.parser:lm-fortigate", "language": "lucene"}},
					{Input: map[string]string{"query": "*", "language": "lucene"}},
				},
				Colors: map[string]string{"failed": "#aaaaaa", "meta.parser:lm-fortigate": "#bbbbbb", "*": "#cccccc"},
			},
		},
		{
			name: "Test params gathering - pie graph with several queries keeps no series filters",
			panel: &lm3.Panel{
				Span: 6, Type: "terms", Chart: "pie", Field: "meta.parser", Size: 10,
				Queries: lm3.Queries{Mode: "all"},
			},
			queries: []lm3.Query{
				{ID: 0, Type: "lucene", Query: "a", Alias: "a", Color: "#aaaaaa"},
				{ID: 1, Type: "lucene", Query: "b", Alias: "b", Color: "#bbbbbb"},
			},
			expected: &visualization.MigrationParams{
				Queries: []lm3.Query{
					{ID: 0, Type: "lucene", Query: "a", Alias: "a", Color: "#aaaaaa"},
					{ID: 1, Type: "lucene", Query: "b", Alias: "b", Color: "#bbbbbb"},
				},
				VisualizationType: &objects.PieGraph{},
				Field:             "meta.parser",
				FieldSize:         10,
				Metric:            "count",
				Span:              6,
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
				Metric:            "count",
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
				Metric:            "count",
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
				Metric:            "count",
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
				Metric:            "count",
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
				Metric:            "count",
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
		{
			name: "Test params gathering - test case 9 - markdown text panel",
			panel: &lm3.Panel{
				Span:    6,
				Type:    "text",
				Mode:    "markdown",
				Content: "# Hello",
				Title:   "Notes",
			},
			expected: &visualization.MigrationParams{
				Title:             "Notes",
				Content:           "# Hello",
				VisualizationType: &objects.Markdown{},
				Metric:            "count",
				Span:              6,
			},
		},
		{
			name: "Test params gathering - test case 10 - html text panel",
			panel: &lm3.Panel{
				Type: "text",
				Mode: "html",
			},
			expectedErr: &PanelTypeNotFoundError{},
		},
		{
			name: "terms_stats total becomes sum of the value field",
			panel: &lm3.Panel{
				Type: "terms", Chart: "bar", Field: "msg.dst_ip.raw", Size: 10, Span: 4,
				Tmode: "terms_stats", Tstat: "total", Valuefield: "msg.rcvd_byte@int.value.raw",
			},
			queries: []lm3.Query{{Type: "lucene", Query: "*"}},
			expected: &visualization.MigrationParams{
				Field: "msg.dst_ip", FieldSize: 10, Span: 4,
				Queries:           []lm3.Query{{Type: "lucene", Query: "*"}},
				VisualizationType: &objects.VerticalGraph{},
				Metric:            "sum", MetricField: "msg.rcvd_byte@int.value",
			},
		},
		{
			name: "terms_stats with unknown tstat falls back to count",
			panel: &lm3.Panel{
				Type: "terms", Chart: "bar", Field: "msg.dst_ip", Size: 10, Span: 4,
				Tmode: "terms_stats", Tstat: "count", Valuefield: "msg.rcvd_byte@int.value",
			},
			queries: []lm3.Query{{Type: "lucene", Query: "*"}},
			expected: &visualization.MigrationParams{
				Field: "msg.dst_ip", FieldSize: 10, Span: 4,
				Queries:           []lm3.Query{{Type: "lucene", Query: "*"}},
				VisualizationType: &objects.VerticalGraph{},
				Metric:            "count",
			},
		},
		{
			name: "histogram max becomes max of the value field",
			panel: &lm3.Panel{
				Type: "histogram", Span: 12, Mode: "max", ValueField: "msg.rcvd_byte@int.value",
			},
			queries: []lm3.Query{{Type: "lucene", Query: "*"}},
			expected: &visualization.MigrationParams{
				Span:              12,
				Queries:           []lm3.Query{{Type: "lucene", Query: "*"}},
				VisualizationType: &objects.EventsOverTime{},
				Metric:            "max", MetricField: "msg.rcvd_byte@int.value",
			},
		},
		{
			name: "plain terms panel with leftover value field stays count",
			panel: &lm3.Panel{
				Type: "terms", Chart: "bar", Span: 12, Tmode: "terms", Tstat: "total", Valuefield: "msg.bytes",
			},
			queries: []lm3.Query{{Type: "lucene", Query: "*"}},
			expected: &visualization.MigrationParams{
				Span:              12,
				Queries:           []lm3.Query{{Type: "lucene", Query: "*"}},
				VisualizationType: &objects.VerticalGraph{},
				Metric:            "count",
			},
		},
		{
			name: "known statistic without value field stays count",
			panel: &lm3.Panel{
				Type: "histogram", Span: 12, Mode: "max", ValueField: "",
			},
			queries: []lm3.Query{{Type: "lucene", Query: "*"}},
			expected: &visualization.MigrationParams{
				Span:              12,
				Queries:           []lm3.Query{{Type: "lucene", Query: "*"}},
				VisualizationType: &objects.EventsOverTime{},
				Metric:            "count",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := gatherMigrationParams(tt.panel, tt.queries, "")
			if tt.expectedErr != nil {
				var target *PanelTypeNotFoundError
				assert.ErrorAs(t, err, &target)
				return
			}
			require.NoError(t, err)

			// Inject ID since it generated inside gather function
			tt.expected.ID = params.ID

			assert.Equal(t, tt.expected, params)
		})
	}
}

func TestMigrate(t *testing.T) {
	tests := []struct {
		name   string
		panels []lm3.Panel
		verify func(t *testing.T, output []lm4.SavedObject)
	}{
		{
			name: "html text panel is skipped and other panels migrate",
			panels: []lm3.Panel{
				{Type: "text", Mode: "html", Title: "Html"},
				{Type: "terms", Chart: "pie", Field: "msg.protocol", Size: 5, Title: "Pie"},
			},
			verify: func(t *testing.T, output []lm4.SavedObject) {
				require.Len(t, output, 2)
				assert.Equal(t, "visualization", output[0].Type)
				assert.Equal(t, "dashboard", output[1].Type)
				require.Len(t, output[1].References, 1)
				assert.Equal(t, "panel_0", output[1].References[0].Name)
			},
		},
		{
			name: "markdown text panel migrates with content and no references",
			panels: []lm3.Panel{
				{Type: "text", Mode: "markdown", Content: "# Hello", Title: "Notes", Span: 6},
			},
			verify: func(t *testing.T, output []lm4.SavedObject) {
				require.Len(t, output, 2)
				assert.Empty(t, output[0].References)

				var visState lm4.VisState
				require.NoError(t, json.Unmarshal([]byte(output[0].Attributes.VisState), &visState))
				assert.Equal(t, "markdown", visState.Type)
				assert.Equal(t, "# Hello", visState.Params.Markdown)
				assert.Contains(t, output[0].Attributes.VisState, `"times":[]`)
				assert.NotContains(t, output[0].Attributes.KibanaSavedObjectMeta.SearchSourceJSON, "indexRefName")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm3Dashboard := &dashboard.LM3Dashboard{
				Title: "Dashboard",
				Rows:  []lm3.Row{{Panels: tt.panels}},
			}

			output, err := Migrate(dashboard.NewLM4Dashboard(), lm3Dashboard, "lm-*")
			require.NoError(t, err)
			tt.verify(t, output)
		})
	}
}
