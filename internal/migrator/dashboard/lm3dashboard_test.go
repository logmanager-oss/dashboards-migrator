package dashboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

func TestNewLM3Dashboard(t *testing.T) {
	input := []byte(`{
		"title": "t",
		"services": {
			"query": {"list": {
				"3": {"id": 3, "query": "c"},
				"1": {"id": 1, "query": "a"},
				"2": {"id": 2, "query": "b"}
			}, "ids": [1, 2, 3]},
			"filter": {"list": {
				"4": {"id": 4, "field": "dst_port", "value": 53},
				"2": {"id": 2, "field": "src", "value": "x"},
				"3": {"id": 3, "field": "note", "value": null},
				"1": {"id": 1, "field": "@timestamp", "value": null}
			}, "ids": [1, 2, 3, 4]}
		},
		"rows": [{"height": "100px", "panels": [{"type": "histogram", "scale": "linear", "sort": null, "value_field": null}]}]
	}`)

	d, err := NewLM3Dashboard(input)
	require.NoError(t, err)

	var queryIDs, filterIDs []int
	for _, q := range d.Queries {
		queryIDs = append(queryIDs, q.ID)
	}
	values := map[int]string{}
	for _, f := range d.GlobalFilters {
		filterIDs = append(filterIDs, f.ID)
		values[f.ID] = string(f.Value)
	}

	assert.Equal(t, []int{1, 2, 3}, queryIDs)
	assert.Equal(t, []int{2, 3, 4}, filterIDs)
	assert.Equal(t, map[int]string{2: "x", 3: "", 4: "53"}, values)
}

func TestNewLM3Dashboard_ObjectFilterValue(t *testing.T) {
	input := []byte(`{"services": {"filter": {"list": {"1": {"id": 1, "field": "f", "value": {"a": 1}}}}}}`)

	_, err := NewLM3Dashboard(input)
	require.Error(t, err)
}

func TestGetPanelQueries(t *testing.T) {
	d := &LM3Dashboard{Queries: []lm3.Query{
		{ID: 0, Query: "a", Pin: true, Enable: true},
		{ID: 1, Query: "b", Enable: true},
		{ID: 2, Query: "c", Pin: true, Enable: true},
		{ID: 3, Query: "d", Enable: false},
	}}

	tests := []struct {
		name     string
		queries  lm3.Queries
		expected []string
	}{
		{name: "all ignores stale ids", queries: lm3.Queries{Mode: "all", IDs: []int{9}}, expected: []string{"a", "b", "c"}},
		{name: "pinned", queries: lm3.Queries{Mode: "pinned"}, expected: []string{"a", "c"}},
		{name: "unpinned drops disabled", queries: lm3.Queries{Mode: "unpinned"}, expected: []string{"b"}},
		{name: "selected", queries: lm3.Queries{Mode: "selected", IDs: []int{1, 2}}, expected: []string{"b", "c"}},
		{name: "selected disabled query is dropped", queries: lm3.Queries{Mode: "selected", IDs: []int{1, 3}}, expected: []string{"b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual []string
			for _, q := range d.GetPanelQueries(&lm3.Panel{Queries: tt.queries}) {
				actual = append(actual, q.Query)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}
