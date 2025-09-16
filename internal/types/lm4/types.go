package lm4

import "time"

type SavedObject struct {
	Attributes       Attributes             `json:"attributes"`
	ID               string                 `json:"id"`
	MigrationVersion map[string]interface{} `json:"migrationVersion"`
	References       []Reference            `json:"references"`
	Type             string                 `json:"type"`
	UpdatedAt        time.Time              `json:"updated_at"`
	Version          string                 `json:"version"`
}

type Attributes struct {
	Description           string                `json:"description"`
	Hits                  *int                  `json:"hits,omitempty"`
	KibanaSavedObjectMeta KibanaSavedObjectMeta `json:"kibanaSavedObjectMeta"`
	OptionsJSON           string                `json:"optionsJSON,omitempty"`
	PanelsJSON            string                `json:"panelsJSON,omitempty"`
	TimeRestore           *bool                 `json:"timeRestore,omitempty"`
	Title                 string                `json:"title"`
	Version               int                   `json:"version"`
	VisState              string                `json:"visState,omitempty"`
	UIStateJSON           string                `json:"uiStateJSON,omitempty"`
	Columns               []string              `json:"columns,omitempty"`
	Sort                  []interface{}         `json:"sort,omitempty"`
}

type Reference struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type KibanaSavedObjectMeta struct {
	SearchSourceJSON string `json:"searchSourceJSON"`
}
