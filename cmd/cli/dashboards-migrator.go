package cli

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/logmanager-oss/dashboards-migrator/internal/config"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/reader"
	"github.com/logmanager-oss/dashboards-migrator/internal/writer"
)

// CLIStart is an entry point for CLI release of dashboard migrator.
// It reads JSON from path privded via flag which is LM3 dashboard and writes output to path provided by flag or to stdout if no output is provided.
func CLIStart() error {
	slog.Info("Starting dashboards migrator...")

	config := &config.Config{}
	config.LoadAndValidate()

	jsonInput, err := reader.ReadFile(config.InputPath)
	if err != nil {
		return fmt.Errorf("reading file: '%s' failed: %v", jsonInput, err)
	}

	var lm3dashboard lm3.Dashboard
	err = json.Unmarshal(jsonInput, &lm3dashboard)
	if err != nil {
		return fmt.Errorf("unmarshalling input: %v", err)
	}

	migrator := migrator.New()

	lm4dashboard, err := migrator.Migrate(lm3dashboard)
	if err != nil {
		return fmt.Errorf("dashboards migration failed: %v", err)
	}

	outputWriter, err := writer.NewWriter(config.OutputPath)
	if err != nil {
		return fmt.Errorf("creating output file: '%s' failed: %v", config.OutputPath, err)
	}

	defer outputWriter.Close()
	defer outputWriter.Flush()

	for _, savedObject := range lm4dashboard {
		rawSavedObject, err := json.Marshal(savedObject)
		if err != nil {
			return fmt.Errorf("marshalling saved object: %v", err)
		}

		err = outputWriter.Write(rawSavedObject)
		if err != nil {
			return fmt.Errorf("writing saved object to file: %v", err)
		}
	}

	return nil
}
