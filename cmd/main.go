package main

import (
	"log/slog"
	"os"

	"github.com/logmanager-oss/dashboards-migrator/cmd/cli"
)

func main() {
	err := cli.CLIStart()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
