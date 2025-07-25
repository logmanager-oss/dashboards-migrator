package lib

import (
	"encoding/json"
	"fmt"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
)

// LIBStart is an entry point for library release of dashboard migrator. It gets byte encoded JSON input which is LM3 dashboard.
// Returns a list of byte encoded output which is an LM4 dashboard. Reason for list usage is due to the fact that LM4 utilises NDJSON for dashboard definition.
// Each line in NDJSON is a seaparate SavedObject (visualisation, search, dashboard or index pattern) which all together form an LM4 dashboard.
func LIBStart(jsonInput []byte) ([][]byte, error) {
	var lm3dashboard lm3.BaseObject
	err := json.Unmarshal(jsonInput, &lm3dashboard)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling input: %v", err)
	}

	migrator := migrator.New()

	lm4dashboard, err := migrator.Migrate(lm3dashboard)
	if err != nil {
		return nil, fmt.Errorf("dashboards migration failed: %v", err)
	}

	var ndjsonOutput [][]byte
	for _, savedObject := range lm4dashboard {
		rawSavedObject, err := json.Marshal(savedObject)
		if err != nil {
			return nil, fmt.Errorf("marshalling output: %v", err)
		}

		ndjsonOutput = append(ndjsonOutput, rawSavedObject)
	}

	return ndjsonOutput, nil
}
