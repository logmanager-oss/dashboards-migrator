package lm3

type BaseObject struct {
	Title    string `json:"title"`
	Services struct {
		Queries struct {
			List map[int]Query `json:"list"`
			IDs  []int         `json:"ids"`
		} `json:"query"`
		GlobalFilters struct {
			List map[int]GlobalFilter `json:"list"`
			IDs  []int                `json:"ids"`
		} `json:"filter"`
	} `json:"services"`
	Rows []Row `json:"rows"`
}

type Query struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Query  string `json:"query"`
	Alias  string `json:"alias"`
	Color  string `json:"color"`
	Pin    bool   `json:"pin"`
	Enable bool   `json:"enable"`
	Field  string `json:"field"`
	Size   int    `json:"size"`
	Union  string `json:"union"`
}

type GlobalFilter struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
	From    string `json:"from"`
	To      string `json:"to"`
	Mandate string `json:"mandate"`
	Active  bool   `json:"active"`
	Alias   string `json:"alias"`
	ID      int    `json:"id"`
	Query   string `json:"query"`
	Value   string `json:"value"`
}

type Row struct {
	Title       string  `json:"title"`
	Height      string  `json:"height"`
	Editable    bool    `json:"editable"`
	Collapse    bool    `json:"collapse"`
	Collapsable bool    `json:"collapsable"`
	Panels      []Panel `json:"panels"`
}

type Panel struct {
	Span        int         `json:"span"`
	Editable    bool        `json:"editable"`
	Group       []string    `json:"group"`
	Type        string      `json:"type"`
	Mode        string      `json:"mode"`
	TimeField   string      `json:"time_field"`
	ValueField  interface{} `json:"value_field"`
	AutoInt     bool        `json:"auto_int"`
	Resolution  int         `json:"resolution"`
	Fill        int         `json:"fill"`
	Linewidth   int         `json:"linewidth"`
	Timezone    string      `json:"timezone"`
	Spyable     bool        `json:"spyable"`
	Zoomlinks   bool        `json:"zoomlinks"`
	Bars        bool        `json:"bars"`
	Stack       bool        `json:"stack"`
	Points      bool        `json:"points"`
	Lines       bool        `json:"lines"`
	Legend      bool        `json:"legend"`
	XAxis       bool        `json:"x-axis"`
	YAxis       bool        `json:"y-axis"`
	Percentage  bool        `json:"percentage"`
	Interactive bool        `json:"interactive"`
	Queries     Queries     `json:"queries"`
	Title       string      `json:"title"`
	Options     bool        `json:"options"`
	Tooltip     struct {
		ValueType    string `json:"value_type"`
		QueryAsAlias bool   `json:"query_as_alias"`
	} `json:"tooltip"`
	Scale   int    `json:"scale"`
	YFormat string `json:"y_format"`
	Grid    struct {
		Max interface{} `json:"max"`
		Min int         `json:"min"`
	} `json:"grid"`
	Annotate struct {
		Enable bool     `json:"enable"`
		Query  string   `json:"query"`
		Size   int      `json:"size"`
		Field  string   `json:"field"`
		Sort   []string `json:"sort"`
	} `json:"annotate"`
	Pointradius  int      `json:"pointradius"`
	ShowQuery    bool     `json:"show_query"`
	LegendCounts bool     `json:"legend_counts"`
	Zerofill     bool     `json:"zerofill"`
	Derivative   bool     `json:"derivative"`
	Interval     string   `json:"interval"`
	Intervals    []string `json:"intervals"`
	Field        string   `json:"field"`
	Size         int      `json:"size"`
	Chart        string   `json:"chart"`
	Fields       []string `json:"fields"`
}

type Queries struct {
	Mode string `json:"mode"`
	IDs  []int  `json:"ids"`
}
