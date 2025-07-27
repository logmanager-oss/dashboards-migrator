package lm3objects

import (
	"encoding/json"
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3"
)

type DashboardObject struct {
	Rows          []lm3.Row
	Filters       []lm3.Query
	GlobalFilters []lm3.Filter
}

func NewDashboard(input []byte) (*DashboardObject, error) {
	var lm3dashboard *lm3.Object
	err := json.Unmarshal(input, &lm3dashboard)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling input: %v", err)
	}

	object := &DashboardObject{}
	object.setQueries(lm3dashboard)
	object.setFilters(lm3dashboard)
	object.setRows(lm3dashboard)

	return object, nil
}

func (d *DashboardObject) setQueries(lm3dashboard *lm3.Object) {
	for _, query := range lm3dashboard.Services.Query.List {
		d.Filters = append(d.Filters, query)
	}
}

func (d *DashboardObject) setFilters(lm3dashboard *lm3.Object) {
	for _, filter := range lm3dashboard.Services.Filter.List {
		d.GlobalFilters = append(d.GlobalFilters, filter)
	}
}

func (d *DashboardObject) setRows(lm3dashboard *lm3.Object) {
	d.Rows = append(d.Rows, lm3dashboard.Rows...)
}
