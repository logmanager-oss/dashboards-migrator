package migrator

import "fmt"

type PanelTypeNotFoundError struct {
	panelType string
	mode      string
}

func (e *PanelTypeNotFoundError) Error() string {
	if e.mode != "" {
		return fmt.Sprintf("panel type [%s/%s] not recognized", e.panelType, e.mode)
	}
	return fmt.Sprintf("panel type [%s] not recognized", e.panelType)
}
