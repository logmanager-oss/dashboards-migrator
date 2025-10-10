package migrator

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func Migrate(lm4Dashboard *dashboard.LM4Dashboard, lm3Dashboard *dashboard.LM3Dashboard) ([]lm4.SavedObject, error) {
	var output []lm4.SavedObject

	for _, row := range lm3Dashboard.Rows {
		for _, panel := range row.Panels {
			params, err := gatherMigrationParams(&panel, lm3Dashboard.GetPanelQueries(&panel))
			if err != nil {
				return nil, fmt.Errorf("building migration params: %w", err)
			}

			visualization, err := visualization.MigratePanelToVisualization(params)
			if err != nil {
				panelTypeNotFoundError := &PanelTypeNotFoundError{}
				if errors.As(err, &panelTypeNotFoundError) {
					slog.Error(err.Error())
					continue
				}
				return nil, fmt.Errorf("migrating %s LM3 panel to LM4 visualization: %w", panel.Title, err)
			}

			lm4Dashboard.LinkVisualizationToDashboardObject(params, visualization.Type)

			output = append(output, *visualization)
		}
	}

	dashboard, err := lm4Dashboard.BuildDashboardObject()
	if err != nil {
		return nil, fmt.Errorf("building LM4 dashboard object: %w", err)
	}

	output = append(output, *dashboard)

	return output, nil
}

func gatherMigrationParams(panel *lm3.Panel, queries []lm3.Query) (*visualization.MigrationParams, error) {
	params := &visualization.MigrationParams{
		ID:      uuid.New().String(),
		Queries: queries,
	}

	var err error
	params.VisualizationType, err = visualisationTypeDiscovery(panel, params.Queries)
	if err != nil {
		return nil, fmt.Errorf("discovering visualization type: %w", err)
	}

	params.Title = panel.Title
	params.Span = panel.Span

	switch params.VisualizationType.(type) {
	case *objects.EventsOverTimeAsSplitSeries:
		// EventsOverTimeAsSplitSeries can have only one query, so this is ok
		params.Field = strings.TrimSuffix(params.Queries[0].Field, ".raw")
		params.FieldSize = params.Queries[0].Size
	case *objects.LogOverview:
		params.Columns = panel.Fields
	default:
		// we no longer use .raw field name convention so we need to strip it
		params.Field = strings.TrimSuffix(panel.Field, ".raw")
		params.FieldSize = panel.Size
	}

	return params, nil
}
