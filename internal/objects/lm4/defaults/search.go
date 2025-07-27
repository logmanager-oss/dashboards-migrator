package defaults

import "github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4"

func GetDefaultSearch() *lm4.SearchSourceJSON {
	return &lm4.SearchSourceJSON{
		Query: lm4.Query{
			Language: "kuery",
			Query:    "",
		},
		Filter: []map[string]interface{}{},
	}
}
