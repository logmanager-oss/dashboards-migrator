package migrator

import (
	"fmt"
	"strings"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualizations"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type Migrator struct {
	lm4Dashboard *dashboard.LM4Dashboard
	lm3Dashboard *dashboard.LM3Dashboard
}

func New(lm4Dashboard *dashboard.LM4Dashboard, lm3Dashboard *dashboard.LM3Dashboard) *Migrator {
	return &Migrator{
		lm4Dashboard: lm4Dashboard,
		lm3Dashboard: lm3Dashboard,
	}
}

func (migrator *Migrator) Migrate(_ string) ([]lm4.SavedObject, error) {
	var output []lm4.SavedObject

	for _, row := range migrator.lm3Dashboard.Rows {
		for _, panel := range row.Panels {
			visualization, err := migrator.migratePanelToVisualization(&panel)
			if err != nil {
				return nil, fmt.Errorf("migrating LM3 panel to LM4 visualization")
			}

			output = append(output, *visualization)
		}
	}

	dashboard, err := migrator.lm4Dashboard.BuildFinalDashboardObject()
	if err != nil {
		return nil, err
	}

	output = append(output, *dashboard)

	return output, nil
}

func (migrator *Migrator) migratePanelToVisualization(panel *lm3.Panel) (*lm4.SavedObject, error) {
	queries := migrator.lm3Dashboard.GetPanelQueries(panel)

	visualizationType, err := migrator.visualisationTypeDiscovery(panel, queries)
	if err != nil {
		return nil, err
	}

	migrationParams := migrator.prepareMigrationParams(panel, queries)

	visualization, err := visualizations.NewLM4Visualization(visualizationType).Migrate(migrationParams)
	if err != nil {
		return nil, fmt.Errorf("migrating %s panel: %v", panel.Title, err)
	}

	return visualization, nil
}

func (migrator *Migrator) prepareMigrationParams(panel *lm3.Panel, queries []lm3.Query) *visualizations.MigrationParams {
	return &visualizations.MigrationParams{
		Title: panel.Title,
		// we no longer use .raw field name convention so we need to strip it
		Field:   strings.TrimSuffix(panel.Field, ".raw"),
		Size:    panel.Size,
		Queries: queries,
	}
}
