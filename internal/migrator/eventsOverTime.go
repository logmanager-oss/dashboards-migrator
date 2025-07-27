package migrator

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/objects"
	vistypes "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/visTypes"
)

func (m *Migrator) migrateEventsOverTimePanel(title string) (*objects.VisualizationObject, error) {
	visualizationObject := objects.NewVisualisation(&vistypes.EventsOverTime{})
	visualizationObject.SetTitle(title)
	visualizationObject.SetSearch(visualizationObject.Type.GetSearch())
	visualizationObject.SetVisStateAggs(visualizationObject.Type.GetAggs())

	m.dashboardObject.SetAndAppendGridData(
		objects.EventsOverTimeVisWidth,
		objects.DefaultVisHeight,
		visualizationObject.ID,
		visualizationObject.Title,
	)

	m.dashboardObject.SetAndAppendReference(
		visualizationObject.ID,
		visualizationObject.RefType,
	)

	return visualizationObject, nil
}
