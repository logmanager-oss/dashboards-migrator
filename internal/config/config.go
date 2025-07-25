package config

import (
	"flag"
	"fmt"
	"os"
)

// Config represents user supplied program input
type Config struct {
	InputPath  string
	OutputPath string
}

// LoadAndValidate loads values from user supplied input into Config struct and validates them
func (c *Config) LoadAndValidate() {
	flag.Func("i", "Path to input file containing LM3 dashboard", c.validateInputPath())

	flag.Func("o", "Path to output file containing LM4 dashboard", c.validateOutputPath())

	flag.Parse()

	// Check if mandatory flags are set
	if c.InputPath == "" {
		fmt.Println("Error: -i flag is mandatory")
		flag.Usage()
		os.Exit(1)
	}
}
