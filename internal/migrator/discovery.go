package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func (m *Migrator) visualisationTypeDiscovery(panel *lm3.Panel) (vistypes.VisType, error) {
	// Histogram panel type means its Events Over Time
	if panel.Type == "histogram" {
		// If there is more then one query, then it must be Events Over Time With Filters
		if len(panel.Queries.IDs) > 1 {
			return &vistypes.EventsOverTimeWithFilters{}, nil
		}
		// Otherwise, if there is only a single query and it is not *, then it also must be Events Over Time With Filters
		if m.lm3Dashboard.Filters[0].Query != "*" {
			return &vistypes.EventsOverTimeWithFilters{}, nil
		}
		// Otherwise its Events Over Time (without filters)
		return &vistypes.EventsOverTime{}, nil
	}

	return nil, fmt.Errorf("not found")
}
