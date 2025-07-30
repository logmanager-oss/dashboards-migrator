package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type Migrator struct {
	savedObjects []lm4.SavedObject
	lm4Dashboard *dashboard.LM4Dashboard
	lm3Dashboard *dashboard.LM3Dashboard
}

func New(lm4Dashboard *dashboard.LM4Dashboard, lm3Dashboard *dashboard.LM3Dashboard) *Migrator {
	return &Migrator{
		lm4Dashboard: lm4Dashboard,
		lm3Dashboard: lm3Dashboard,
	}
}

func (m *Migrator) Migrate() ([]lm4.SavedObject, error) {
	for _, row := range m.lm3Dashboard.Rows {
		for _, panel := range row.Panels {
			filters := m.lm3Dashboard.GetPanelFilters(&panel)

			visualisationType, err := m.visualisationTypeDiscovery(&panel, filters)
			if err != nil {
				return nil, err
			}

			visualization, err := visualization.NewLM4Visualisation(
				panel.Title,
				filters,
				visualisationType,
			)
			if err != nil {
				return nil, fmt.Errorf("migrating %s panel: %v", panel.Title, err)
			}

			m.appendRefAndGridToDashboard(visualization)
			m.appendSavedObjectToOutput(visualization)
		}
	}

	finalDashboardObject, err := m.lm4Dashboard.BuildFinalDashboardObject()
	if err != nil {
		return nil, err
	}

	m.appendSavedObjectToOutput(finalDashboardObject)

	return m.savedObjects, nil
}

func (m *Migrator) appendSavedObjectToOutput(savedObject *lm4.SavedObject) {
	m.savedObjects = append(m.savedObjects, *savedObject)
}

func (m *Migrator) appendRefAndGridToDashboard(visualizationObject *lm4.SavedObject) {
	m.lm4Dashboard.SetAndAppendGridData(
		visualization.EventsOverTimeVisWidth,
		visualization.DefaultVisHeight,
		visualizationObject.ID,
		visualizationObject.Attributes.Title,
	)

	m.lm4Dashboard.SetAndAppendReference(
		visualizationObject.ID,
		visualizationObject.Type,
	)
}
