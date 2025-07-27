package migrator

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3/lm3objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/lm4objects"
)

type Migrator struct {
	savedObjects []lm4.SavedObject
}

func New(*lm4objects.DashboardObject, *lm3objects.DashboardObject) *Migrator {
	return &Migrator{}
}

func (m *Migrator) Migrate() ([]lm4.SavedObject, error) {
	return m.savedObjects, nil
}
