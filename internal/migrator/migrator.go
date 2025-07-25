package migrator

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type Migrator struct {
	savedObjects []lm4.SavedObject
}

func New() *Migrator {
	return &Migrator{}
}

func (m *Migrator) Migrate(_ lm3.BaseObject) ([]lm4.SavedObject, error) {
	return m.savedObjects, nil
}
