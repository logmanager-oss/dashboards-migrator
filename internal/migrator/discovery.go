package migrator

import (
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualizations/vistypes"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func (migrator *Migrator) visualisationTypeDiscovery(panel *lm3.Panel, queries []lm3.Query) (vistypes.VisType, error) { // nolint
	// TODO: add a logic that will recognize LM4 visualization type from LM3 panel type
	return nil, fmt.Errorf("not found")
}
