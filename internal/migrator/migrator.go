package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/objects"
)

type Migrator struct {
	savedObjects    []lm4.SavedObject
	dashboardObject *objects.DashboardObject
}

func New() *Migrator {
	return &Migrator{
		dashboardObject: objects.NewDashboard(),
	}
}

func (m *Migrator) Migrate(lm3Dashboard lm3.Dashboard) ([]lm4.SavedObject, error) {
	for _, row := range lm3Dashboard.Rows {
		for _, panel := range row.Panels {
			if panel.Type == "histogram" {
				if panel.Queries.Mode == "all" {
					newVisualizationObject, err := m.migrateEventsOverTimePanel(panel.Title)
					if err != nil {
						return nil, fmt.Errorf("migrating %s panel: %v", panel.Title, err)
					}

					finalVisualizationObject, err := newVisualizationObject.GetFinalVisualizationObject()
					if err != nil {
						return nil, err
					}

					m.appendSavedObjectToOutput(finalVisualizationObject)
				}
			}
		}
	}

	finalDashboardObject, err := m.dashboardObject.GetFinalDashboardObject()
	if err != nil {
		return nil, err
	}

	m.appendSavedObjectToOutput(finalDashboardObject)

	return m.savedObjects, nil
}

func (m *Migrator) appendSavedObjectToOutput(savedObject *lm4.SavedObject) {
	m.savedObjects = append(m.savedObjects, *savedObject)
}
