package visualizations

import (
	"encoding/json"
	"testing"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualizations/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

	"github.com/stretchr/testify/assert"
)

func TestMigrator_migrateVisualizations(t *testing.T) {
	tests := []struct {
		name              string
		title             string
		visualizationType vistypes.VisType
		queries           []lm3.Query
		queriesIDs        []int
		field             string
		size              int
		expected          string
		columns           []string
	}{
		{
			name:              "Test case: migrate panel: events over time",
			title:             "Events Over Time",
			visualizationType: &vistypes.EventsOverTime{},
			queries: []lm3.Query{
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
			name:              "Test case: migrate panel: events over time with queries",
			title:             "Events Over Time With Queries",
			visualizationType: &vistypes.EventsOverTime{},
			queries: []lm3.Query{
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
					Query:  "meta.tags:loginfailed",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"meta.tags:loginsuccess or meta.tags:loginfailed\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time With Queries","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time With Queries\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"field\":\"@timestamp\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":100,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: events over time As Split Series",
			title:             "Events Over Time As Split Series",
			visualizationType: &vistypes.EventsOverTimeAsSplitSeries{},
			queries: []lm3.Query{
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
			field:    "meta.parser",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time As Split Series","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time As Split Series\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"params\":{\"field\":\"@timestamp\",\"timeRange\":{\"from\":\"now-15m\",\"to\":\"now\"},\"useNormalizedOpenSearchInterval\":true,\"scaleMetricValues\":true,\"interval\":\"auto\",\"drop_partials\":false,\"min_doc_count\":1,\"extended_bounds\":{}},\"schema\":\"segment\"},{\"id\":\"3\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.parser\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"group\"}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newLM4visualization := NewLM4Visualization(tt.visualizationType)

			migrationParams := &MigrationParams{
				tt.title,
				tt.field,
				tt.size,
				tt.queries,
				tt.columns,
			}

			actualSavedObject, err := newLM4visualization.Migrate(migrationParams)
			if err != nil {
				t.Error(err)
			}

			var expectedSavedObject *lm4.SavedObject
			err = json.Unmarshal([]byte(tt.expected), &expectedSavedObject)
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

			assert.Equal(t, expectedSearch, actualSearch)

			// Comparation of saved objects happens without nested visState and KibanaSavedObjectMeta objects since we are unable to predict order of keys in marshalled structs.
			// As such those objects are compared separately (above)
			actualSavedObject.Attributes.VisState = ""
			expectedSavedObject.Attributes.VisState = ""
			actualSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = ""
			expectedSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = ""
			expectedSavedObject.ID = actualSavedObject.ID

			assert.Equal(t, expectedSavedObject, actualSavedObject)
		})
	}
}
