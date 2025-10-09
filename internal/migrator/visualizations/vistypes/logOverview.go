package vistypes

import (
	"time"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type LogOverview struct{}

func (lo *LogOverview) GetDefaultVisualizationSavedObject() *lm4.SavedObject {
	return &lm4.SavedObject{
		Attributes: lm4.Attributes{
			Columns:     nil,
			Description: "",
			Hits:        new(int),
			KibanaSavedObjectMeta: lm4.KibanaSavedObjectMeta{
				SearchSourceJSON: "",
			},
			Sort:    []interface{}{},
			Title:   "",
			Version: 1,
		},
		ID:               "",
		MigrationVersion: map[string]interface{}{"search": "7.9.3"},
		References: []lm4.Reference{
			{
				ID:   "",
				Name: "kibanaSavedObjectMeta.searchSourceJSON.index",
				Type: "index-pattern",
			},
		},
		Type:      "search",
		UpdatedAt: time.Time{},
		Version:   "",
	}
}

func (lo *LogOverview) GetVisualizationConfig(string, int) []lm4.VisStateAggs {
	return []lm4.VisStateAggs{}
}

func (lo *LogOverview) GetDefaultVisState() *lm4.VisState {
	return &lm4.VisState{}
}
