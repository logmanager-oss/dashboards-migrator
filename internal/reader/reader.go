// Package reader provides functionality for reading input files.
package reader

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	cleanPath := filepath.Clean(path)

	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
