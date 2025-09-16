package cli

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/logmanager-oss/dashboards-migrator/cmd/config"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/dashboard"
	"github.com/logmanager-oss/dashboards-migrator/internal/reader"
	"github.com/logmanager-oss/dashboards-migrator/internal/writer"
)

// CLIStart is an entry point for CLI release of dashboard migrator.
// It reads JSON from path provided via flag which is LM3 dashboard and writes output to path provided by flag or to stdout if no output is provided.
func CLIStart() error {
	slog.Info("Starting dashboards migrator...")

	config := &config.Config{}
	err := config.LoadAndValidate()
	if err != nil {
		return fmt.Errorf("parsing flags: %v", err)
	}

	jsonInput, err := reader.ReadFile(config.InputPath)
	if err != nil {
		return fmt.Errorf("reading file: '%s' failed: %v", jsonInput, err)
	}

	lm3Dashboard, err := dashboard.NewLM3Dashboard(jsonInput)
	if err != nil {
		return err
	}

	outputWriter, err := writer.NewWriter(config.OutputPath)
	if err != nil {
		return fmt.Errorf("creating output file: '%s' failed: %v", config.OutputPath, err)
	}

	defer outputWriter.Close()
	defer outputWriter.Flush()

	lm4Dashboard := dashboard.NewLM4Dashboard()

	migrator := migrator.New(lm4Dashboard, lm3Dashboard)
	lm4dashboard, err := migrator.Migrate(config.IndexPattern)
	if err != nil {
		return fmt.Errorf("dashboards migration failed: %v", err)
	}

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

	slog.Info("Migration finished!")

	return nil
}
