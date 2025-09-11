package migrator

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
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

func (migrator *Migrator) Migrate(_ string) ([]lm4.SavedObject, error) {
	return migrator.savedObjects, nil
}
