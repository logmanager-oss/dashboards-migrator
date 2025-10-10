package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func visualisationTypeDiscovery(panel *lm3.Panel, queries []lm3.Query) (objects.VisType, error) { // nolint
	// Histogram panel type means its Events Over Time
	if panel.Type == "histogram" {
		// If there is only a single filter and it is topN, then it must be Events Over Time As Split Series
		if len(queries) == 1 && queries[0].Type == "topN" {
			return &objects.EventsOverTimeAsSplitSeries{}, nil
		}

		return &objects.EventsOverTime{}, nil
	}

	if panel.Type == "table" {
		return &objects.LogOverview{}, nil
	}

	if panel.Type == "map" {
		return &objects.Map{}, nil
	}

	if panel.Type == "terms" {
		if panel.Chart == "bar" {
			return &objects.VerticalGraph{}, nil
		}
		if panel.Chart == "pie" {
			return &objects.PieGraph{}, nil
		}
		if panel.Chart == "table" {
			return &objects.TableGraph{}, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
