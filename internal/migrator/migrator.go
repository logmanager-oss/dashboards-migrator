package migrator

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"
)

type Migrator struct {
	savedObjects []lm4.SavedObject
}

func New() *Migrator {
	return &Migrator{}
}

func (m *Migrator) Migrate(_ lm3.Dashboard) ([]lm4.SavedObject, error) {
	return m.savedObjects, nil
}
