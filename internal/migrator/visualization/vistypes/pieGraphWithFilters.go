package vistypes

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type PieGrapWithFilters struct{}

func (pg *PieGrapWithFilters) GetVisualizationConfig(filters []lm4.Filter, field string, size int) []lm4.VisStateAggs {
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
			ID:      "3",
			Enabled: true,
			Type:    "filters",
			Schema:  "split",
			Params: lm4.VisStateAggsParams{
				Filters: filters,
			},
		},
		{
			ID:      "4",
			Enabled: true,
			Type:    "terms",
			Schema:  "segment",
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
	}
}

func (pg *PieGrapWithFilters) GetDefaultVisState() *lm4.VisState {
	return defaults.GetDefaultPieWithFiltersGraphVisState()
}
