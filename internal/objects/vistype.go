package objects

import (
	"cmp"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

// AggConfig carries the LM3 panel values that LM4 aggregations are built from.
type AggConfig struct {
	Field       string
	Size        int
	Metric      string
	MetricField string
	Filters     []lm4.Filter
}

type VisType interface {
	GetDefaultVisualizationSavedObject(string) *lm4.SavedObject
	GetVisualizationConfig(AggConfig) []lm4.VisStateAggs
	GetDefaultVisState() *lm4.VisState
}

func metricAgg(cfg AggConfig) lm4.VisStateAggs {
	return lm4.VisStateAggs{
		ID:      "1",
		Enabled: true,
		Type:    cmp.Or(cfg.Metric, "count"),
		Params:  lm4.VisStateAggsParams{Field: cfg.MetricField},
		Schema:  "metric",
	}
}

// MetricLabel returns the series label that OpenSearch Dashboards generates for the metric.
func MetricLabel(cfg AggConfig) string {
	switch cfg.Metric {
	case "sum":
		return "Sum of " + cfg.MetricField
	case "avg":
		return "Average " + cfg.MetricField
	case "min":
		return "Min " + cfg.MetricField
	case "max":
		return "Max " + cfg.MetricField
	default:
		return "Count"
	}
}
