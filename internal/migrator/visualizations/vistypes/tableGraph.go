package vistypes

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type TableGraph struct{}

func (tg *TableGraph) GetDefaultVisualizationSavedObject() *lm4.SavedObject {
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

func (tg *TableGraph) GetVisualizationConfig(field string, size int) []lm4.VisStateAggs {
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
				OtherBucket:        true,
				OtherBucketLabel:   "Other",
				MissingBucket:      false,
				MissingBucketLabel: "Missing",
			},
			Schema: "bucket",
		},
	}
}

func (tg *TableGraph) GetDefaultVisState() *lm4.VisState {
	return &lm4.VisState{
		Title: "",
		Type:  "table",
		Aggs:  []lm4.VisStateAggs{},
		Params: lm4.VisStateParams{
			PerPage:                10,
			ShowPartialRows:        false,
			ShowMetricsAtAllLevels: false,
			ShowTotal:              false,
			TotalFunc:              "sum",
			PercentageCol:          "",
		},
	}
}
