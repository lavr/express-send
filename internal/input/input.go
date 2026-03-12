package input

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadMessage reads the message text from one of:
// 1. --message-from flag (file path)
// 2. Positional argument
// 3. stdin (if not a terminal)
func ReadMessage(messageFrom string, args []string, stdin io.Reader, isTerminal bool) (string, error) {
	if messageFrom != "" {
		data, err := os.ReadFile(messageFrom)
		if err != nil {
			return "", fmt.Errorf("reading message from %q: %w", messageFrom, err)
		}
		return strings.TrimRight(string(data), "\n"), nil
	}

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	if isTerminal {
		return "", fmt.Errorf("no message provided: use argument, --message-from, or pipe via stdin")
	}

	data, err := io.ReadAll(stdin)
	if err != nil {
		return "", fmt.Errorf("reading stdin: %w", err)
	}

	msg := strings.TrimRight(string(data), "\n")
	if msg == "" {
		return "", fmt.Errorf("empty message from stdin")
	}
	return msg, nil
}
