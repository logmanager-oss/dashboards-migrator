package vistypes

import "github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

type VisType interface {
	GetVisualizationConfig([]lm4.Filter, string, int) []lm4.VisStateAggs
}
