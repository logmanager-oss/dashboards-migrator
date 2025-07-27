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
}

type Reference struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type KibanaSavedObjectMeta struct {
	SearchSourceJSON string
}

type SearchSourceJSON struct {
	Query        Query                    `json:"query"`
	Filter       []map[string]interface{} `json:"filter"`
	IndexRefName string                   `json:"indexRefName,omitempty"`
}

type Query struct {
	Language string `json:"language"`
	Query    string `json:"query"`
}

type PanelJSON struct {
	Version          string   `json:"version"`
	GridData         GridData `json:"gridData"`
	PanelIndex       string   `json:"panelIndex"`
	EmbeddableConfig struct {
		HidePanelTitles bool `json:"hidePanelTitles"`
	} `json:"embeddableConfig"`
	Title        string `json:"title"`
	PanelRefName string `json:"panelRefName"`
}

type GridData struct {
	X int    `json:"x"`
	Y int    `json:"y"`
	W int    `json:"w"`
	H int    `json:"h"`
	I string `json:"i"`
}

type VisState struct {
	Title  string         `json:"title"`
	Type   string         `json:"type"`
	Aggs   []VisStateAggs `json:"aggs"`
	Params VisStateParams `json:"params"`
}

type VisStateAggs struct {
	ID      string             `json:"id"`
	Enabled bool               `json:"enabled"`
	Type    string             `json:"type"`
	Schema  string             `json:"schema"`
	Params  VisStateAggsParams `json:"params"`
}

type VisStateAggsParams struct {
	Field                           string            `json:"field,omitempty"`
	OrderBy                         string            `json:"orderBy,omitempty"`
	Order                           string            `json:"order,omitempty"`
	Size                            int               `json:"size,omitempty"`
	OtherBucket                     bool              `json:"otherBucket,omitempty"`
	OtherBucketLabel                string            `json:"otherBucketLabel,omitempty"`
	MissingBucket                   bool              `json:"missingBucket,omitempty"`
	MissingBucketLabel              string            `json:"missingBucketLabel,omitempty"`
	Filters                         []Filter          `json:"filters,omitempty"`
	TimeRange                       map[string]string `json:"timeRange,omitempty"`
	UseNormalizedOpenSearchInterval bool              `json:"useNormalizedOpenSearchInterval,omitempty"`
	ScaleMetricValues               bool              `json:"scaleMetricValues,omitempty"`
	Interval                        string            `json:"interval,omitempty"`
	DropPartials                    bool              `json:"drop_partials,omitempty"`
	MinDocCount                     int               `json:"min_doc_count,omitempty"`
	ExtendedBounds                  struct {
	} `json:"extended_bounds,omitempty"`
}

type Filter struct {
	Input map[string]string `json:"input"`
	Label string            `json:"label"`
}

type VisStateParams struct {
	Type           string                   `json:"type"`
	Grid           map[string]interface{}   `json:"grid"`
	CategoryAxes   []map[string]interface{} `json:"categoryAxes"`
	ValueAxes      []map[string]interface{} `json:"valueAxes"`
	SeriesParams   []map[string]interface{} `json:"seriesParams"`
	AddTooltip     bool                     `json:"addTooltip"`
	AddLegend      bool                     `json:"addLegend"`
	LegendPosition string                   `json:"legendPosition"`
	Times          []interface{}            `json:"times"`
	AddTimeMarker  bool                     `json:"addTimeMarker"`
	Labels         map[string]interface{}   `json:"labels"`
	ThresholdLine  map[string]interface{}   `json:"thresholdLine"`
}
