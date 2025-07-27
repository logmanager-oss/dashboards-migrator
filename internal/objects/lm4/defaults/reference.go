package defaults

import "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"

func GetDefaultReference() lm4.Reference {
	return lm4.Reference{
		ID:   "",
		Name: "",
		Type: "",
	}
}
