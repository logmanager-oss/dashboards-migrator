package objects

import "github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

func GetDefaultPanelObject() *lm4.PanelJSON {
	return &lm4.PanelJSON{
		Version: "2.19.1",
		GridData: lm4.GridData{
			X: 0,
			Y: 0,
			W: 0,
			H: 0,
			I: "",
		},
		PanelIndex: "",
		EmbeddableConfig: lm4.EmbeddableConfig{
			HidePanelTitles: false,
		},
		Title:        "",
		PanelRefName: "",
	}
}
