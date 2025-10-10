package objects

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type Map struct{}

func (m *Map) GetDefaultVisualizationSavedObject(indexPattern string) *lm4.SavedObject {
	return &lm4.SavedObject{
		Attributes: lm4.Attributes{
			Description: "",
			KibanaSavedObjectMeta: lm4.KibanaSavedObjectMeta{
				SearchSourceJSON: "",
			},
			Title:       "",
			UIStateJSON: "{}",
			Version:     1,
			VisState:    "",
		},
		ID:               "",
		MigrationVersion: map[string]interface{}{"visualization": "7.10.0"},
		References: []lm4.Reference{
			{
				ID:   indexPattern,
				Name: "kibanaSavedObjectMeta.searchSourceJSON.index",
				Type: "index-pattern",
			},
		},
		Type:      "visualization",
		UpdatedAt: time.Time{},
		Version:   "",
	}
}

func (m *Map) GetVisualizationConfig(field string, size int) []lm4.VisStateAggs {
	return []lm4.VisStateAggs{
		{
			ID:      "1",
			Enabled: true,
			Type:    "count",
			Params: lm4.VisStateAggsParams{
				Filters: []lm4.Filter{},
			},
			Schema: "metric",
		},
		{
			ID:      "2",
			Enabled: true,
			Type:    "terms",
			Params: lm4.VisStateAggsParams{
				Field:              field,
				OrderBy:            "1",
				Order:              "desc",
				Size:               size,
				OtherBucket:        false,
				OtherBucketLabel:   "Other",
				MissingBucket:      false,
				MissingBucketLabel: "Missing",
				Exclude:            "Un",
			},
			Schema: "segment",
		},
	}
}

func (m *Map) GetDefaultVisState() *lm4.VisState {
	return &lm4.VisState{
		Title: "",
		Type:  "region_map",
		Aggs:  []lm4.VisStateAggs{},
		Params: lm4.VisStateParams{
			LayerChosenByUser: "default",
			LegendPosition:    "bottomright",
			AddTooltip:        true,
			ColorSchema:       "Yellow to Red",
			EmsHotLink:        "",
			IsDisplayWarning:  true,
			Wms: map[string]interface{}{
				"enabled": false,
				"url":     "",
				"options": map[string]interface{}{
					"version":     "",
					"layers":      "",
					"format":      "image/png",
					"transparent": true,
					"attribution": "",
					"styles":      "",
				},
				"selectedTmsLayer": map[string]interface{}{
					"origin":      "elastic_maps_service",
					"id":          "road_map",
					"minZoom":     0,
					"maxZoom":     22,
					"attribution": "<a rel=\"noreferrer noopener\" href=\"https://www.openstreetmap.org/copyright\">Map data © OpenStreetMap contributors</a>",
				},
			},
			MapZoom: 2,
			MapCenter: []int{
				0,
				0,
			},
			OutlineWeight: 1,
			ShowAllShapes: true,
			SelectedLayer: map[string]interface{}{
				"name": "planet",
				"url":  "/vendor/maps/ne_50m_admin_0_countries.geojson?v=1",
				"meta": map[string]interface{}{
					"feature_collection_path": "features",
				},
				"attribution": "Custom GeoJSON – Local",
				"fields": []map[string]interface{}{
					{
						"name":        "ISO_A2",
						"description": "ISO Alpha-2 Country Code",
					},
				},
				"format": map[string]interface{}{
					"type": "geojson",
				},
				"layerId": "self_hosted.planet",
				"isEMS":   false,
			},
			SelectedJoinField: map[string]interface{}{
				"name":        "ISO_A2",
				"description": "ISO Alpha-2 Country Code",
			},
			SelectedCustomLayer: map[string]interface{}{
				"name": "planet",
				"url":  "/vendor/maps/ne_50m_admin_0_countries.geojson?v=1",
				"meta": map[string]interface{}{
					"feature_collection_path": "features",
				},
				"attribution": "Custom GeoJSON – Local",
				"fields": []map[string]interface{}{
					{
						"name":        "ISO_A2",
						"description": "ISO Alpha-2 Country Code",
					},
				},
				"format": map[string]interface{}{
					"type": "geojson",
				},
				"layerId": "self_hosted.planet",
				"isEMS":   false,
			},
			SelectedCustomJoinField: map[string]interface{}{
				"name":        "ISO_A2",
				"description": "ISO Alpha-2 Country Code",
			},
		},
	}
}
