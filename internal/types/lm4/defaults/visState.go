package defaults

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func GetDefaultHistogramVisState() *lm4.VisState { // nolint:dupl // we don't care about duplicates here
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

func GetDefaultVerticalGraphVisState() *lm4.VisState { // nolint:dupl // we don't care about duplicates here
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
					"position": "top",
					"show":     false,
					"style":    map[string]interface{}{},
					"scale": map[string]interface{}{
						"type": "linear",
					},
					"labels": map[string]interface{}{
						"show":     true,
						"filter":   true,
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
					"mode": "normal",
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

func GetDefaultVerticalGraphWithFiltersVisState() *lm4.VisState { // nolint:dupl // we don't care about duplicates here
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
						"filter":   true,
						"rotate":   0,
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
					"mode": "normal",
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
