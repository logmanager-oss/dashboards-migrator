package migrator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
)

func TestMigrator_migrateEventsOverTimePanel(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		expected string
	}{
		{
			name:     "Test case: migrate events over time panel",
			title:    "Events Over Time",
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Events Over Time","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Events Over Time\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"date_histogram\",\"schema\":\"segment\",\"params\":{\"field\":\"@timestamp\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":100,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"bottom\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":false,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"stacked\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migrator := New()

			actual, err := migrator.migrateEventsOverTimePanel(tt.title)
			if err != nil {
				t.Error(err)
			}

			var expectedSavedObject *lm4.SavedObject
			err = json.Unmarshal([]byte(tt.expected), &expectedSavedObject)
			if err != nil {
				t.Error(err)
			}

			var expectedVisState *lm4.VisState
			err = json.Unmarshal([]byte(expectedSavedObject.Attributes.VisState), &expectedVisState)
			if err != nil {
				t.Error(err)
			}

			var expectedSearch *lm4.SearchSourceJSON
			err = json.Unmarshal([]byte(expectedSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON), &expectedSearch)
			if err != nil {
				t.Error(err)
			}

			expectedSavedObject.Attributes.VisState = ""
			expectedSavedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = ""

			assert.Equal(t, expectedSavedObject, actual.SavedObject)
			assert.Equal(t, expectedVisState, actual.VisState)
			assert.Equal(t, expectedSearch, actual.Search)
		})
	}
}
