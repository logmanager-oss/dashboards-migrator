package objects

import "github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"

func GetDefaultSearchObject(withRef bool) *lm4.SearchSourceJSON {
	search := &lm4.SearchSourceJSON{
		Query: lm4.Query{
			Language: "kuery",
			Query:    "",
		},
		Filter: []lm4.GlobalFilter{},
	}

	if withRef {
		search.IndexRefName = "kibanaSavedObjectMeta.searchSourceJSON.index"
	}

	return search
}
