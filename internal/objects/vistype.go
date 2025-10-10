package objects

import "github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

type VisType interface {
	GetDefaultVisualizationSavedObject() *lm4.SavedObject
	GetVisualizationConfig(string, int) []lm4.VisStateAggs
	GetDefaultVisState() *lm4.VisState
}
