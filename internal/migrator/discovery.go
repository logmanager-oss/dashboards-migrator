package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func (m *Migrator) visualisationTypeDiscovery(panel *lm3.Panel, filters []lm3.Filter) (vistypes.VisType, error) { // nolint
	// Histogram panel type means its Events Over Time
	if panel.Type == "histogram" {
		if len(filters) == 1 {
			// If there is only a single filter and it is *, then it must be Events Over Time
			if filters[0].Query == "*" {
				return &vistypes.EventsOverTime{}, nil
			}
			// If there is only a single filter and it is topN, then it must be Events Over Time As Split Series
			if filters[0].Type == "topN" {
				return &vistypes.EventsOverTimeAsSplitSeries{}, nil
			}
			// Otherwise it must be Events Over Time With Filters
			return &vistypes.EventsOverTimeWithFilters{}, nil
		}
		// If there is more then one filter, then it must be Events Over Time With Filters
		if len(filters) > 1 {
			return &vistypes.EventsOverTimeWithFilters{}, nil
		}
	}

	if panel.Type == "terms" {
		if len(filters) == 1 {
			// If there is only a single filter and it is *, then it must be Vertical Or Pie graph
			if filters[0].Query == "*" {
				if panel.Chart == "bar" {
					return &vistypes.VerticalGraph{}, nil
				}
				if panel.Chart == "pie" {
					return &vistypes.PieGraph{}, nil
				}
			}
			// Otherwise it will be Vertical Or Pie Graph With Filters
			if panel.Chart == "bar" {
				return &vistypes.VerticalGraphWithFilters{}, nil
			}
			if panel.Chart == "pie" {
				return &vistypes.PieGrapWithFilters{}, nil
			}
		}
		// If there is more then one filter, then it must be Vertical Or Pie Graph With Filters
		if len(filters) > 1 {
			if panel.Chart == "bar" {
				return &vistypes.VerticalGraphWithFilters{}, nil
			}
			if panel.Chart == "pie" {
				return &vistypes.PieGrapWithFilters{}, nil
			}
		}
	}

	return nil, fmt.Errorf("not found")
}
