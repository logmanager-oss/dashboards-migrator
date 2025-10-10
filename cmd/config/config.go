package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Config represents user supplied program input
type Config struct {
	InputPath    string
	OutputPath   string
	IndexPattern string
}

// LoadAndValidate loads values from user supplied input into Config struct and validates them
func (c *Config) LoadAndValidate() error {
	flag.StringVar(&c.InputPath, "i", "", "Path to input file containing LM3 dashboard (Mandatory)")
	flag.StringVar(&c.OutputPath, "o", "", "Path to output file containing LM4 dashboard (Mandatory)")
	flag.StringVar(&c.IndexPattern, "ip", "", "Index-pattern string (Default: empty)")

	flag.Parse()

	err := c.validateInputPath(c.InputPath)
	if err != nil {
		return fmt.Errorf("validating input path: %w", err)
	}

	err = c.validateOutputPath(c.OutputPath)
	if err != nil {
		return fmt.Errorf("validating output path: %w", err)
	}

	err = c.validateIndexPattern(c.IndexPattern)
	if err != nil {
		return fmt.Errorf("validating index-pattern: %w", err)
	}

	return nil
}

func (c *Config) validateInputPath(flagValue string) error {
	if flagValue == "" {
		return fmt.Errorf("path to input file (-i flag) cannot be empty")
	}

	fileInfo, err := os.Stat(flagValue)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("input %s cannot be a directory", flagValue)
	}

	return nil
}

func (c *Config) validateOutputPath(flagValue string) error {
	if flagValue == "" {
		return fmt.Errorf("path to output file (-o flag) cannot be empty")
	}

	fileInfo, err := os.Stat(flagValue)
	if err != nil {
		// If output path does not exist it's ok - we will create it
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	// If output path exists check if it's a directory - which would be wrong
	if fileInfo.IsDir() {
		return fmt.Errorf("output file %s cannot be a directory", flagValue)
	}

	return nil
}

func (c *Config) validateIndexPattern(flagValue string) error {
	if !strings.HasPrefix(flagValue, "lm-") {
		return fmt.Errorf("index-pattern must begin with \"lm-\" prefix")
	}

	return nil
}
