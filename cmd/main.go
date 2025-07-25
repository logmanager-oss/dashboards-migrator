package main

import (
	"log/slog"

	"github.com/logmanager-oss/dashboards-migrator/cmd/cli"
)

func main() {
	err := cli.CLIStart()
	if err != nil {
		slog.Error(err.Error())
	}
}
