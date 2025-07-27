package cli

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/logmanager-oss/dashboards-migrator/internal/config"
	"github.com/logmanager-oss/dashboards-migrator/internal/migrator"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm3/lm3objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects/lm4/lm4objects"
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

	lm4DashboardObject := lm4objects.NewDashboard()
	lm3DashboardObject, err := lm3objects.NewDashboard(jsonInput)
	if err != nil {
		return err
	}

	migrator := migrator.New(lm4DashboardObject, lm3DashboardObject)
	lm4dashboard, err := migrator.Migrate()
	if err != nil {
		return fmt.Errorf("dashboards migration failed: %v", err)
	}

	outputWriter, err := writer.NewWriter(config.OutputPath)
	if err != nil {
		return fmt.Errorf("creating output file: '%s' failed: %v", config.OutputPath, err)
	}

	defer outputWriter.Flush()
	defer outputWriter.Close()

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
