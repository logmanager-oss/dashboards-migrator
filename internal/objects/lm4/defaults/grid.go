package defaults

import "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"

func GetDefaultGridData() *lm4.PanelJSON {
	return &lm4.PanelJSON{
		Version: "2.19.1",
		GridData: struct {
			X int    "json:\"x\""
			Y int    "json:\"y\""
			W int    "json:\"w\""
			H int    "json:\"h\""
			I string "json:\"i\""
		}{
			X: 0,
			Y: 0,
			W: 0,
			H: 0,
			I: "",
		},
		PanelIndex: "",
		EmbeddableConfig: struct {
			HidePanelTitles bool "json:\"hidePanelTitles\""
		}{
			HidePanelTitles: false,
		},
		Title:        "",
		PanelRefName: "",
	}
}
