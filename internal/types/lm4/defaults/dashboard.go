package defaults

import (
	"time"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func GetDefaultDashboardSavedObject() *lm4.SavedObject {
	return &lm4.SavedObject{
		Attributes: lm4.Attributes{
			Description: "",
			Hits:        new(int),
			KibanaSavedObjectMeta: lm4.KibanaSavedObjectMeta{
				SearchSourceJSON: "",
			},
			OptionsJSON: "{\"hidePanelTitles\":false,\"useMargins\":true}",
			PanelsJSON:  "",
			TimeRestore: new(bool),
			Title:       "",
			Version:     1,
		},
		ID:               uuid.New().String(),
		MigrationVersion: map[string]interface{}{"dashboard": "7.9.3"},
		References:       nil,
		Type:             "dashboard",
		UpdatedAt:        time.Time{},
		Version:          "",
	}
}
