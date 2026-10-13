// Package migrator orchestrates the migration of dashboards from LM3 to LM4.
package migrator

import (
	"cmp"
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

func Migrate(lm4Dashboard *dashboard.LM4Dashboard, lm3Dashboard *dashboard.LM3Dashboard, indexPattern string) ([]lm4.SavedObject, error) {
	var output []lm4.SavedObject

	for _, row := range lm3Dashboard.Rows {
		for _, panel := range row.Panels {
			params, err := gatherMigrationParams(&panel, lm3Dashboard.GetPanelQueries(&panel), indexPattern)
			if err != nil {
				if errors.As(err, new(*PanelTypeNotFoundError)) {
					slog.Warn("skipping panel", "title", panel.Title, "type", panel.Type, "mode", panel.Mode)
					continue
				}
				return nil, fmt.Errorf("building migration params: %w", err)
			}

			visualization, err := visualization.MigratePanelToVisualization(params)
			if err != nil {
				return nil, fmt.Errorf("migrating %s LM3 panel to LM4 visualization: %w", panel.Title, err)
			}

			lm4Dashboard.LinkVisualizationToDashboardObject(params, visualization.Type)

			output = append(output, *visualization)
		}
	}

	dashboard, err := lm4Dashboard.MigrateDashboard(lm3Dashboard.Title, lm3Dashboard.GlobalFilters, indexPattern)
	if err != nil {
		return nil, fmt.Errorf("building LM4 dashboard object: %w", err)
	}

	output = append(output, *dashboard)

	return output, nil
}

func gatherMigrationParams(panel *lm3.Panel, queries []lm3.Query, indexPattern string) (*visualization.MigrationParams, error) {
	params := &visualization.MigrationParams{
		ID:           uuid.New().String(),
		Queries:      queries,
		IndexPattern: indexPattern,
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
		params.Field = visualization.StripRaw(params.Queries[0].Field)
		params.FieldSize = params.Queries[0].Size
	case *objects.LogOverview:
		params.Columns = panel.Fields
	case *objects.Markdown:
		params.Content = panel.Content
	default:
		params.Field = visualization.StripRaw(panel.Field)
		params.FieldSize = panel.Size
	}

	if _, ok := params.VisualizationType.(*objects.EventsOverTime); ok && len(queries) > 1 {
		params.Filters, params.Colors = gatherSeriesFilters(queries)
	}

	params.Metric, params.MetricField = gatherMetric(panel)

	return params, nil
}

// gatherSeriesFilters builds one filters-agg entry per query, so LM4 draws one series per query like LM3.
func gatherSeriesFilters(queries []lm3.Query) ([]lm4.Filter, map[string]string) {
	filters := make([]lm4.Filter, 0, len(queries))
	colors := make(map[string]string, len(queries))
	for _, q := range queries {
		// An empty query means all events, which the filters agg spells "*".
		query := cmp.Or(visualization.NormalizeLucene(strings.TrimSpace(q.Query)), "*")
		filters = append(filters, lm4.Filter{
			Input: map[string]string{"query": query, "language": "lucene"},
			Label: q.Alias,
		})
		// OSD keys series colors by the filter label, or by the query text when the label is empty.
		if q.Color != "" {
			colors[cmp.Or(q.Alias, query)] = q.Color
		}
	}

	return filters, colors
}

var _metricTypes = map[string]string{"total": "sum", "mean": "avg", "min": "min", "max": "max"}

// gatherMetric maps the LM3 value-field statistic of a panel to an LM4 metric aggregation.
// Panels without a usable statistic and value field stay on count.
func gatherMetric(panel *lm3.Panel) (metric, field string) {
	var stat, valueField string
	switch {
	case panel.Type == "terms" && panel.Tmode == "terms_stats":
		stat, valueField = panel.Tstat, panel.Valuefield
	case panel.Type == "histogram":
		stat, valueField = panel.Mode, panel.ValueField
	}

	metric, ok := _metricTypes[stat]
	if !ok || valueField == "" {
		return "count", ""
	}

	return metric, visualization.StripRaw(valueField)
}
