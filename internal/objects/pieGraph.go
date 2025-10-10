package objects

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type PieGraph struct{}

func (pg *PieGraph) GetDefaultVisualizationSavedObject() *lm4.SavedObject {
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

func (pg *PieGraph) GetVisualizationConfig(field string, size int) []lm4.VisStateAggs {
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
	return &lm4.VisState{
		Title: "",
		Type:  "pie",
		Aggs:  []lm4.VisStateAggs{},
		Params: lm4.VisStateParams{
			Type:           "pie",
			AddTooltip:     true,
			AddLegend:      true,
			LegendPosition: "right",
			IsDonut:        false,
			Labels: map[string]interface{}{
				"show":       true,
				"values":     true,
				"last_level": true,
				"truncate":   100,
			},
		},
	}
}
