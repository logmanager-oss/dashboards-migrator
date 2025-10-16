package migrator

import "fmt"

type PanelTypeNotFoundError struct {
	panelType string
}

func (e *PanelTypeNotFoundError) Error() string {
	return fmt.Sprintf("panel type [%s] not recognized", e.panelType)
}
