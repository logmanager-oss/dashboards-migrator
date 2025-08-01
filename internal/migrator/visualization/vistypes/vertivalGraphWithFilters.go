package vistypes

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type VerticalGraphWithFilters struct{}

func (e *VerticalGraphWithFilters) GetVisualizationConfig(filters []lm4.Filter, field string, size int) []lm4.VisStateAggs {
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
				Filters:            []lm4.Filter{},
			},
		},
		{
			ID:      "3",
			Enabled: true,
			Type:    "filters",
			Schema:  "segment",
			Params: lm4.VisStateAggsParams{
				Filters: filters,
			},
		},
	}
}

func (e *VerticalGraphWithFilters) GetDefaultVisState() *lm4.VisState {
	return defaults.GetDefaultVerticalGraphWithFiltersVisState()
}
