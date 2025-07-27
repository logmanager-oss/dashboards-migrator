package vistypes

import "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"

type VisType interface {
	GetAggs() []lm4.VisStateAggs
	GetSearch() *lm4.SearchSourceJSON
}
