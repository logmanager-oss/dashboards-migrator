package dashboard

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

type LM3Dashboard struct {
	Rows          []lm3.Row
	Queries       []lm3.Query
	GlobalFilters []lm3.GlobalFilter
}

func NewLM3Dashboard(input []byte) (*LM3Dashboard, error) {
	var baseLM3Object *lm3.BaseObject
	err := json.Unmarshal(input, &baseLM3Object)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling input: %v", err)
	}

	lm3dashboard := &LM3Dashboard{}
	lm3dashboard.unpackQueries(baseLM3Object)
	lm3dashboard.unpackFilters(baseLM3Object)
	lm3dashboard.unpackRows(baseLM3Object)

	return lm3dashboard, nil
}

func (d *LM3Dashboard) unpackQueries(lm3dashboard *lm3.BaseObject) {
	for _, query := range lm3dashboard.Services.Queries.List {
		d.Queries = append(d.Queries, query)
	}
}

func (d *LM3Dashboard) unpackFilters(lm3dashboard *lm3.BaseObject) {
	for _, filter := range lm3dashboard.Services.GlobalFilters.List {
		d.GlobalFilters = append(d.GlobalFilters, filter)
	}
}

func (d *LM3Dashboard) unpackRows(lm3dashboard *lm3.BaseObject) {
	d.Rows = append(d.Rows, lm3dashboard.Rows...)
}

func (d *LM3Dashboard) GetPanelQueries(panel *lm3.Panel) []lm3.Query {
	var queries []lm3.Query
	for _, query := range d.Queries {
		if slices.Contains(panel.Queries.IDs, query.ID) {
			queries = append(queries, query)
		}
	}

	return queries
}
