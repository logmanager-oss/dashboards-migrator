package vistypes

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/defaults"
)

type EventsOverTime struct{}

func (e *EventsOverTime) GetAggs() []lm4.VisStateAggs {
	return []lm4.VisStateAggs{
		{
			ID:      "1",
			Enabled: true,
			Type:    "count",
			Params:  lm4.VisStateAggsParams{},
			Schema:  "metric",
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
			},
		},
	}
}

func (e *EventsOverTime) GetSearch() *lm4.SearchSourceJSON {
	search := defaults.GetDefaultSearch()
	search.IndexRefName = "kibanaSavedObjectMeta.searchSourceJSON.index"

	return search
}
