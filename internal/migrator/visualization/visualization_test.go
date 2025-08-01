package visualization

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func TestMigrator_migrateEventsOverTimePanel(t *testing.T) {
	tests := []struct {
		name              string
		title             string
		visualizationType vistypes.VisType
		filters           []lm3.Filter
		filterIDs         []int
		field             string
		size              int
		expected          string
	}{
		{
			name:              "Test case: migrate panel: events over time",
			title:             "Events Over Time",
			visualizationType: &vistypes.EventsOverTime{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "*",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"field\":\"@timestamp\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":100,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: events over time with filters",
			title:             "Events Over Time With Filters",
			visualizationType: &vistypes.EventsOverTimeWithFilters{},
			filters: []lm3.Filter{
				{
					ID:     1,
					Type:   "lucene",
					Query:  "meta.tags:loginsuccess",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     0,
					Type:   "lucene",
					Query:  "*",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time With Filters\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"params\":{\"field\":\"@timestamp\",\"timeRange\":{\"from\":\"now-15m\",\"to\":\"now\"},\"useNormalizedOpenSearchInterval\":true,\"scaleMetricValues\":true,\"interval\":\"auto\",\"drop_partials\":false,\"min_doc_count\":1,\"extended_bounds\":{}},\"schema\":\"segment\"},{\"id\":\"3\",\"enabled\":true,\"type\":\"filters\",\"params\":{\"filters\":[{\"input\":{\"query\":\"meta.tags:loginsuccess\",\"language\":\"kuery\"},\"label\":\"\"},{\"input\":{\"query\":\"*\",\"language\":\"kuery\"},\"label\":\"\"}]},\"schema\":\"group\"}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: events over time As Split Series",
			title:             "Events Over Time As Split Series",
			visualizationType: &vistypes.EventsOverTimeAsSplitSeries{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "topN",
					Query:  "*",
					Field:  "meta.parser",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
					Size:   10,
					Union:  "AND",
				},
			},
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time As Split Series","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time As Split Series\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"params\":{\"field\":\"@timestamp\",\"timeRange\":{\"from\":\"now-15m\",\"to\":\"now\"},\"useNormalizedOpenSearchInterval\":true,\"scaleMetricValues\":true,\"interval\":\"auto\",\"drop_partials\":false,\"min_doc_count\":1,\"extended_bounds\":{}},\"schema\":\"segment\"},{\"id\":\"3\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.parser\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"group\"}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: vertical graph",
			title:             "Vertical Graph",
			visualizationType: &vistypes.VerticalGraph{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "*",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "msg.command",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Vertical Graph","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Vertical Graph\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"msg.command\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"top\",\"show\":false,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":true,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"normal\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: vertical graph",
			title:             "Vertical Graph With Filters",
			visualizationType: &vistypes.VerticalGraphWithFilters{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "meta.parser:lm-fortigate",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     1,
					Type:   "lucene",
					Query:  "meta.parser:lm-hp-comware",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "meta.src.severity",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Vertical Graph With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Vertical Graph With Filters\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.src.severity\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"group\"},{\"id\":\"3\",\"enabled\":true,\"type\":\"filters\",\"params\":{\"filters\":[{\"input\":{\"query\":\"meta.parser:lm-fortigate\",\"language\":\"kuery\"},\"label\":\"\"},{\"input\":{\"query\":\"meta.parser:lm-hp-comware\",\"language\":\"kuery\"},\"label\":\"\"}]},\"schema\":\"segment\"}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":true,\"truncate\":100,\"rotate\":0},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"normal\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"},\"orderBucketsBySum\":false}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: pie graph",
			title:             "Pie Graph",
			visualizationType: &vistypes.PieGraph{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "*",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "meta.src.severity",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Pie Graph","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Pie Graph\",\"type\":\"pie\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.src.severity\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"segment\"}],\"params\":{\"type\":\"pie\",\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"right\",\"isDonut\":false,\"labels\":{\"show\":true,\"values\":true,\"last_level\":true,\"truncate\":100}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: pie graph with filters",
			title:             "Pie Graph With Filters",
			visualizationType: &vistypes.PieGrapWithFilters{},
			filters: []lm3.Filter{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "meta.parser:lm-fortigate",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     1,
					Type:   "lucene",
					Query:  "meta.parser:lm-hp-comware",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "meta.src.severity",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Pie Graph With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Pie Graph With Filters\",\"type\":\"pie\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"3\",\"enabled\":true,\"type\":\"filters\",\"params\":{\"filters\":[{\"input\":{\"query\":\"meta.parser:lm-fortigate\",\"language\":\"kuery\"},\"label\":\"\"},{\"input\":{\"query\":\"meta.parser:lm-hp-comware\",\"language\":\"kuery\"},\"label\":\"\"}]},\"schema\":\"split\"},{\"id\":\"4\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.src.severity\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"segment\"}],\"params\":{\"type\":\"pie\",\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"right\",\"isDonut\":false,\"labels\":{\"show\":true,\"values\":true,\"last_level\":true,\"truncate\":100},\"row\":false}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSavedObject, err := NewLM4Visualisation(
				tt.title,
				tt.filters,
				tt.field,
				tt.size,
				tt.visualizationType,
			)
			if err != nil {
				t.Error(err)
			}

			var expectedSavedObject *lm4.SavedObject
			err = json.Unmarshal([]byte(tt.expected), &expectedSavedObject)
			if err != nil {
				t.Error(err)
			}

			var actualVisState *lm4.VisState
			err = json.Unmarshal([]byte(actualSavedObject.Attributes.VisState), &actualVisState)
			if err != nil {
				t.Error(err)
			}

			var expectedVisState *lm4.VisState
			err = json.Unmarshal([]byte(expectedSavedObject.Attributes.VisState), &expectedVisState)
			if err != nil {
				t.Error(err)
			}

			var actualSearch *lm4.SearchSourceJSON
			err = json.Unmarshal([]byte(actualSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON), &actualSearch)
			if err != nil {
				t.Error(err)
			}

			var expectedSearch *lm4.SearchSourceJSON
			err = json.Unmarshal([]byte(expectedSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON), &expectedSearch)
			if err != nil {
				t.Error(err)
			}

			actualSavedObject.Attributes.VisState = ""
			expectedSavedObject.Attributes.VisState = ""
			actualSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = ""
			expectedSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = ""

			assert.Equal(t, expectedSavedObject, actualSavedObject)
			assert.Equal(t, expectedVisState, actualVisState)
			assert.Equal(t, expectedSearch, actualSearch)
		})
	}
}
