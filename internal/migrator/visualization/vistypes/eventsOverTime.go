package vistypes

import "github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

type EventsOverTime struct{}

func (e *EventsOverTime) GetVisualizationConfig([]lm4.Filter, string, int) []lm4.VisStateAggs {
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
				Field:              "@timestamp",
				OrderBy:            "1",
				Order:              "desc",
				Size:               100,
				OtherBucket:        true,
				OtherBucketLabel:   "Other",
				MissingBucket:      false,
				MissingBucketLabel: "Missing",
				Filters:            []lm4.Filter{},
			},
		},
	}
}
