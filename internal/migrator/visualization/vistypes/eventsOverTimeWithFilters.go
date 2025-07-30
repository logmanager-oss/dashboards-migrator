package vistypes

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type EventsOverTimeWithFilters struct{}

func (e *EventsOverTimeWithFilters) GetVisualizationConfig(filters []lm4.Filter) []lm4.VisStateAggs {
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
			Type:    "filters",
			Schema:  "group",
			Params: lm4.VisStateAggsParams{
				Filters: filters,
			},
		},
	}
}
