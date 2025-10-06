package writer

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

type OutputWriter struct {
	buffer *bufio.Writer
	f      *os.File
}

func NewWriter(path string) (*OutputWriter, error) {
	var outputFile *os.File
	var err error
	if path != "" {
		outputFile, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("opening output file for writing: %w", err)
		}

	} else {
		slog.Debug("output path empty - defaulting to stdout")
		outputFile = os.Stdout
	}

	outputWriter := bufio.NewWriter(outputFile)

	return &OutputWriter{buffer: outputWriter, f: outputFile}, nil
}

func (w *OutputWriter) Write(p []byte) error {
	_, err := fmt.Fprintln(w.buffer, string(p))
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}

func (w *OutputWriter) Close() {
	if closeErr := w.f.Close(); closeErr != nil {
		fmt.Printf("warning: failed to close output file '%s': %v\n", w.f.Name(), closeErr)
	}
}

func (w *OutputWriter) Flush() {
	if flushErr := w.buffer.Flush(); flushErr != nil {
		fmt.Printf("warning: failed to flush output buffer: %v\n", flushErr)
	}
}
