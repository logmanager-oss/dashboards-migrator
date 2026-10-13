// Package lm3 provides type definitions for Logmanager 3 dashboard objects.
package lm3

import (
	"encoding/json"
	"fmt"
)

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
}

// FlexString decodes a JSON string, number or bool into its text form.
// JSON null decodes to "".
type FlexString string

func (f *FlexString) UnmarshalJSON(data []byte) error {
	switch data[0] {
	case 'n':
		*f = ""
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("unmarshalling string value: %w", err)
		}
		*f = FlexString(s)
	case '{', '[':
		return fmt.Errorf("unsupported value %s", data)
	default:
		*f = FlexString(data)
	}
	return nil
}

type GlobalFilter struct {
	Type    string     `json:"type"`
	Field   string     `json:"field"`
	Mandate string     `json:"mandate"`
	Active  bool       `json:"active"`
	Alias   string     `json:"alias"`
	ID      int        `json:"id"`
	Query   string     `json:"query"`
	Value   FlexString `json:"value"`
}

type Row struct {
	Height string  `json:"height"`
	Panels []Panel `json:"panels"`
}

type Panel struct {
	Span       int      `json:"span"`
	Type       string   `json:"type"`
	Mode       string   `json:"mode"`
	Content    string   `json:"content"`
	Chart      string   `json:"chart"`
	Field      string   `json:"field"`
	Size       int      `json:"size"`
	Fields     []string `json:"fields"`
	Sort       []string `json:"sort"`
	Title      string   `json:"title"`
	Queries    Queries  `json:"queries"`
	Tmode      string   `json:"tmode"`
	Tstat      string   `json:"tstat"`
	Valuefield string   `json:"valuefield"`
	ValueField string   `json:"value_field"`
}

type Queries struct {
	Mode string `json:"mode"`
	IDs  []int  `json:"ids"`
}
