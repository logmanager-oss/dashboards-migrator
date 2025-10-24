package objects

import (
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

func GetDefaultFilterObject(indexRefName string) *lm4.GlobalFilter {
	filter := &lm4.GlobalFilter{
		Meta: lm4.GlobalFilterMeta{
			Alias:        nil,
			Negate:       false,
			Disabled:     false,
			Type:         "phrase",
			Key:          "",
			IndexRefName: indexRefName,
		},
		State: map[string]string{"store": "appState"},
	}

	return filter
}
