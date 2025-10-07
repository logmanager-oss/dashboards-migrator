package vistypes

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type EventsOverTimeAsSplitSeries struct{}

func (e *EventsOverTimeAsSplitSeries) GetDefaultVisualizationSavedObject() *lm4.SavedObject {
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
				// TODO: id should be set to user defined index pattern
				ID:   "",
				Name: "kibanaSavedObjectMeta.searchSourceJSON.index",
				Type: "index-pattern",
			},
		},
		Type:      "visualization",
		UpdatedAt: time.Time{},
		Version:   "",
	}
}

func (e *EventsOverTimeAsSplitSeries) GetVisualizationConfig(field string, size int) []lm4.VisStateAggs {
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
			Type:    "date_histogram",
			Schema:  "segment",
			Params: lm4.VisStateAggsParams{
				Field: "@timestamp",
				TimeRange: map[string]string{
					"from": "now-15m",
					"to":   "now",
				},
				UseNormalizedOpenSearchInterval: true,
				ScaleMetricValues:               true,
				DropPartials:                    false,
				Interval:                        "auto",
				MinDocCount:                     1,
				ExtendedBounds:                  struct{}{},
			},
		},
		{
			ID:      "3",
			Enabled: true,
			Type:    "terms",
			Schema:  "group",
			Params: lm4.VisStateAggsParams{
				Field:              field,
				OrderBy:            "1",
				Order:              "desc",
				Size:               size,
				OtherBucket:        true,
				OtherBucketLabel:   "Other",
				MissingBucket:      false,
				MissingBucketLabel: "Missing",
			},
		},
	}
}

func (e *EventsOverTimeAsSplitSeries) GetDefaultVisState() *lm4.VisState { // nolint:dupl
	return &lm4.VisState{
		Title: "",
		Type:  "histogram",
		Aggs:  []lm4.VisStateAggs{},
		Params: lm4.VisStateParams{
			Type: "histogram",
			Grid: map[string]interface{}{
				"categoryLines": true,
				"valueAxis":     "ValueAxis-1",
			},
			CategoryAxes: []map[string]interface{}{
				{
					"id":       "CategoryAxis-1",
					"type":     "category",
					"position": "bottom",
					"show":     true,
					"style":    map[string]interface{}{},
					"scale": map[string]interface{}{
						"type": "linear",
					},
					"labels": map[string]interface{}{
						"show":     true,
						"filter":   false,
						"truncate": float64(100),
					},
					"title": map[string]interface{}{},
				},
			},
			ValueAxes: []map[string]interface{}{
				{
					"id":       "ValueAxis-1",
					"name":     "LeftAxis-1",
					"type":     "value",
					"position": "left",
					"show":     true,
					"style":    map[string]interface{}{},
					"scale": map[string]interface{}{
						"type": "linear",
						"mode": "normal",
					},
					"labels": map[string]interface{}{
						"show":     true,
						"rotate":   float64(0),
						"filter":   false,
						"truncate": float64(100),
					},
					"title": map[string]interface{}{
						"text": "Count",
					},
				},
			},
			SeriesParams: []map[string]interface{}{
				{
					"show": true,
					"type": "histogram",
					"mode": "stacked",
					"data": map[string]interface{}{
						"label": "Count",
						"id":    "1",
					},
					"valueAxis":              "ValueAxis-1",
					"drawLinesBetweenPoints": true,
					"lineWidth":              float64(2),
					"showCircles":            true,
				},
			},
			AddTooltip:     true,
			AddLegend:      true,
			LegendPosition: "top",
			Times:          []interface{}{},
			AddTimeMarker:  false,
			Labels:         map[string]interface{}{"show": true},
			ThresholdLine: map[string]interface{}{
				"show":  false,
				"value": float64(10),
				"width": float64(1),
				"style": "full",
				"color": "#E7664C",
			},
		},
	}
}
