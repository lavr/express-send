package cmd

import (
	"encoding/json"
	"io"
)

// printOutput writes data as JSON (with indent) if format is "json",
// otherwise calls textFn for human-readable output.
func printOutput(w io.Writer, format string, textFn func(), data any) error {
	if format == "json" {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(data)
	}
	textFn()
	return nil
}
