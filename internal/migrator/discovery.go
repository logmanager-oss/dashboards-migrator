package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualizations/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func (migrator *Migrator) visualisationTypeDiscovery(panel *lm3.Panel, queries []lm3.Query) (vistypes.VisType, error) { // nolint
	// Histogram panel type means its Events Over Time
	if panel.Type == "histogram" {
		// If there is only a single filter and it is topN, then it must be Events Over Time As Split Series
		if len(queries) == 1 && queries[0].Type == "topN" {
			return &vistypes.EventsOverTimeAsSplitSeries{}, nil
		}

		return &vistypes.EventsOverTime{}, nil
	}

	if panel.Type == "table" {
		return &vistypes.LogOverview{}, nil
	}

	if panel.Type == "map" {
		return &vistypes.Map{}, nil
	}

	return nil, fmt.Errorf("not found")
}
