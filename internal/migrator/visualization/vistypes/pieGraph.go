package vistypes

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4/defaults"
)

type PieGraph struct{}

func (pg *PieGraph) GetVisualizationConfig(_ []lm4.Filter, field string, size int) []lm4.VisStateAggs {
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

func (pg *PieGraph) GetDefaultVisState() *lm4.VisState {
	return defaults.GetDefaultPieGraphVisState()
}
