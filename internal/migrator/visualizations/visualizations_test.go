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
		{
			name:              "Test case: migrate panel: log overview",
			title:             "Log Overview",
			visualizationType: &vistypes.LogOverview{},
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
			columns: []string{
				"meta.src.ip@ip.value",
				"raw",
			},
			expected: `{"attributes":{"columns":["meta.src.ip@ip.value","raw"],"description":"","hits":0,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"highlightAll\":true,\"version\":true,\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"sort":[],"title":"Log Overview","version":1},"id":"","migrationVersion":{"search":"7.9.3"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"search","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: log overview with filters",
			title:             "Log Overview With Filters",
			visualizationType: &vistypes.LogOverview{},
			queries: []lm3.Query{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "msg.protocol:TCP",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     1,
					Type:   "lucene",
					Query:  "msg.protocol:UDP",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			columns: []string{
				"meta.src.ip@ip.value",
				"raw",
			},
			expected: `{"attributes":{"columns":["meta.src.ip@ip.value","raw"],"description":"","hits":0,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"highlightAll\":true,\"version\":true,\"query\":{\"query\":\"msg.protocol:TCP or msg.protocol:UDP\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"sort":[],"title":"Log Overview With Filters","version":1},"id":"","migrationVersion":{"search":"7.9.3"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"search","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: map",
			title:             "Map",
			visualizationType: &vistypes.Map{},
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
			field:    "msg.dst_ip@ip.country_code",
			size:     100,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Map","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Map\",\"type\":\"region_map\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"msg.dst_ip@ip.country_code\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":100,\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"exclude\":\"Un\"},\"schema\":\"segment\"}],\"params\":{\"layerChosenByUser\":\"default\",\"legendPosition\":\"bottomright\",\"addTooltip\":true,\"colorSchema\":\"Yellow to Red\",\"emsHotLink\":\"\",\"isDisplayWarning\":true,\"wms\":{\"enabled\":false,\"url\":\"\",\"options\":{\"version\":\"\",\"layers\":\"\",\"format\":\"image/png\",\"transparent\":true,\"attribution\":\"\",\"styles\":\"\"},\"selectedTmsLayer\":{\"origin\":\"elastic_maps_service\",\"id\":\"road_map\",\"minZoom\":0,\"maxZoom\":22,\"attribution\":\"<a rel=\\\"noreferrer noopener\\\" href=\\\"https://www.openstreetmap.org/copyright\\\">Map data © OpenStreetMap contributors</a>\"}},\"mapZoom\":2,\"mapCenter\":[0,0],\"outlineWeight\":1,\"showAllShapes\":true,\"selectedLayer\":{\"name\":\"planet\",\"url\":\"/vendor/maps/ne_50m_admin_0_countries.geojson?v=1\",\"meta\":{\"feature_collection_path\":\"features\"},\"attribution\":\"Custom GeoJSON – Local\",\"fields\":[{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}],\"format\":{\"type\":\"geojson\"},\"layerId\":\"self_hosted.planet\",\"isEMS\":false},\"selectedJoinField\":{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"},\"selectedCustomLayer\":{\"name\":\"planet\",\"url\":\"/vendor/maps/ne_50m_admin_0_countries.geojson?v=1\",\"meta\":{\"feature_collection_path\":\"features\"},\"attribution\":\"Custom GeoJSON – Local\",\"fields\":[{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}],\"format\":{\"type\":\"geojson\"},\"layerId\":\"self_hosted.planet\",\"isEMS\":false},\"selectedCustomJoinField\":{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: map with filters",
			title:             "Map With Filters",
			visualizationType: &vistypes.Map{},
			queries: []lm3.Query{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "msg.dst_ip@ip.country_code:US",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     1,
					Type:   "lucene",
					Query:  "msg.dst_ip@ip.country_code:CZ",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "msg.dst_ip@ip.country_code",
			size:     100,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"msg.dst_ip@ip.country_code:US or msg.dst_ip@ip.country_code:CZ\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Map With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Map With Filters\",\"type\":\"region_map\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"msg.dst_ip@ip.country_code\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":100,\"otherBucket\":false,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\",\"exclude\":\"Un\"},\"schema\":\"segment\"}],\"params\":{\"layerChosenByUser\":\"default\",\"legendPosition\":\"bottomright\",\"addTooltip\":true,\"colorSchema\":\"Yellow to Red\",\"emsHotLink\":\"\",\"isDisplayWarning\":true,\"wms\":{\"enabled\":false,\"url\":\"\",\"options\":{\"version\":\"\",\"layers\":\"\",\"format\":\"image/png\",\"transparent\":true,\"attribution\":\"\",\"styles\":\"\"},\"selectedTmsLayer\":{\"origin\":\"elastic_maps_service\",\"id\":\"road_map\",\"minZoom\":0,\"maxZoom\":22,\"attribution\":\"<a rel=\\\"noreferrer noopener\\\" href=\\\"https://www.openstreetmap.org/copyright\\\">Map data © OpenStreetMap contributors</a>\"}},\"mapZoom\":2,\"mapCenter\":[0,0],\"outlineWeight\":1,\"showAllShapes\":true,\"selectedLayer\":{\"name\":\"planet\",\"url\":\"/vendor/maps/ne_50m_admin_0_countries.geojson?v=1\",\"meta\":{\"feature_collection_path\":\"features\"},\"attribution\":\"Custom GeoJSON – Local\",\"fields\":[{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}],\"format\":{\"type\":\"geojson\"},\"layerId\":\"self_hosted.planet\",\"isEMS\":false},\"selectedJoinField\":{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"},\"selectedCustomLayer\":{\"name\":\"planet\",\"url\":\"/vendor/maps/ne_50m_admin_0_countries.geojson?v=1\",\"meta\":{\"feature_collection_path\":\"features\"},\"attribution\":\"Custom GeoJSON – Local\",\"fields\":[{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}],\"format\":{\"type\":\"geojson\"},\"layerId\":\"self_hosted.planet\",\"isEMS\":false},\"selectedCustomJoinField\":{\"name\":\"ISO_A2\",\"description\":\"ISO Alpha-2 Country Code\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: vertical graph",
			title:             "Vertical Graph",
			visualizationType: &vistypes.VerticalGraph{},
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
			field:    "msg.command",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Vertical Graph","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Vertical Graph\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"msg.command\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"top\",\"show\":false,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":true,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"normal\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: vertical graph",
			title:             "Vertical Graph With Filters",
			visualizationType: &vistypes.VerticalGraph{},
			queries: []lm3.Query{
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
			field:    "msg.command",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"meta.parser:lm-fortigate or meta.parser:lm-hp-comware\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Vertical Graph With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Vertical Graph With Filters\",\"type\":\"histogram\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"schema\":\"metric\",\"params\":{\"field\":\"\",\"orderBy\":\"\",\"order\":\"\",\"size\":0,\"otherBucket\":false,\"otherBucketLabel\":\"\",\"missingBucket\":false,\"missingBucketLabel\":\"\"}},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"schema\":\"group\",\"params\":{\"field\":\"msg.command\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"}}],\"params\":{\"type\":\"histogram\",\"grid\":{\"categoryLines\":true,\"valueAxis\":\"ValueAxis-1\"},\"categoryAxes\":[{\"id\":\"CategoryAxis-1\",\"type\":\"category\",\"position\":\"top\",\"show\":false,\"style\":{},\"scale\":{\"type\":\"linear\"},\"labels\":{\"show\":true,\"filter\":true,\"truncate\":100},\"title\":{}}],\"valueAxes\":[{\"id\":\"ValueAxis-1\",\"name\":\"LeftAxis-1\",\"type\":\"value\",\"position\":\"left\",\"show\":true,\"style\":{},\"scale\":{\"type\":\"linear\",\"mode\":\"normal\"},\"labels\":{\"show\":true,\"rotate\":0,\"filter\":false,\"truncate\":100},\"title\":{\"text\":\"Count\"}}],\"seriesParams\":[{\"show\":true,\"type\":\"histogram\",\"mode\":\"normal\",\"data\":{\"label\":\"Count\",\"id\":\"1\"},\"valueAxis\":\"ValueAxis-1\",\"drawLinesBetweenPoints\":true,\"lineWidth\":2,\"showCircles\":true}],\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"top\",\"times\":[],\"addTimeMarker\":false,\"labels\":{\"show\":true},\"thresholdLine\":{\"show\":false,\"value\":10,\"width\":1,\"style\":\"full\",\"color\":\"#E7664C\"}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: pie graph",
			title:             "Pie Graph",
			visualizationType: &vistypes.PieGraph{},
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
			field:    "meta.src.severity",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Pie Graph","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Pie Graph\",\"type\":\"pie\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.src.severity\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"segment\"}],\"params\":{\"type\":\"pie\",\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"right\",\"isDonut\":false,\"labels\":{\"show\":true,\"values\":true,\"last_level\":true,\"truncate\":100}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: pie graph with filters",
			title:             "Pie Graph With Filters",
			visualizationType: &vistypes.PieGraph{},
			queries: []lm3.Query{
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
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"meta.parser:lm-fortigate or meta.parser:lm-hp-comware\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Pie Graph With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Pie Graph With Filters\",\"type\":\"pie\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"meta.src.severity\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"segment\"}],\"params\":{\"type\":\"pie\",\"addTooltip\":true,\"addLegend\":true,\"legendPosition\":\"right\",\"isDonut\":false,\"labels\":{\"show\":true,\"values\":true,\"last_level\":true,\"truncate\":100}}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: table graph",
			title:             "Table Graph",
			visualizationType: &vistypes.TableGraph{},
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
			field:    "msg.protocol",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Table Graph","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Table Graph\",\"type\":\"table\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"msg.protocol\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"bucket\"}],\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMetricsAtAllLevels\":false,\"showTotal\":false,\"totalFunc\":\"sum\",\"percentageCol\":\"\"}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
		},
		{
			name:              "Test case: migrate panel: table graph with filters",
			title:             "Table Graph With Filters",
			visualizationType: &vistypes.TableGraph{},
			queries: []lm3.Query{
				{
					ID:     0,
					Type:   "lucene",
					Query:  "msg.protocol:TCP",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
				{
					ID:     1,
					Type:   "lucene",
					Query:  "msg.protocol:UDP",
					Alias:  "",
					Color:  "",
					Pin:    false,
					Enable: true,
				},
			},
			field:    "msg.protocol",
			size:     10,
			expected: `{"attributes":{"description":"","kibanaSavedObjectMeta":{"searchSourceJSON":"{\"query\":{\"query\":\"msg.protocol:TCP or msg.protocol:UDP\",\"language\":\"kuery\"},\"filter\":[],\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"},"title":"Table Graph With Filters","uiStateJSON":"{}","version":1,"visState":"{\"title\":\"Table Graph With Filters\",\"type\":\"table\",\"aggs\":[{\"id\":\"1\",\"enabled\":true,\"type\":\"count\",\"params\":{},\"schema\":\"metric\"},{\"id\":\"2\",\"enabled\":true,\"type\":\"terms\",\"params\":{\"field\":\"msg.protocol\",\"orderBy\":\"1\",\"order\":\"desc\",\"size\":10,\"otherBucket\":true,\"otherBucketLabel\":\"Other\",\"missingBucket\":false,\"missingBucketLabel\":\"Missing\"},\"schema\":\"bucket\"}],\"params\":{\"perPage\":10,\"showPartialRows\":false,\"showMetricsAtAllLevels\":false,\"showTotal\":false,\"totalFunc\":\"sum\",\"percentageCol\":\"\"}}"},"id":"","migrationVersion":{"visualization":"7.10.0"},"references":[{"id":"","name":"kibanaSavedObjectMeta.searchSourceJSON.index","type":"index-pattern"}],"type":"visualization","updated_at":"0001-01-01T00:00:00Z","version":""}`,
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
