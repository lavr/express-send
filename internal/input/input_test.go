package input

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadMessage_FromArg(t *testing.T) {
	msg, err := ReadMessage("", []string{"hello world"}, nil, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "hello world" {
		t.Errorf("got %q, want %q", msg, "hello world")
	}
}

func TestReadMessage_FromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "msg.txt")
	if err := os.WriteFile(path, []byte("file content\n"), 0644); err != nil {
		t.Fatal(err)
	}

	msg, err := ReadMessage(path, nil, nil, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "file content" {
		t.Errorf("got %q, want %q", msg, "file content")
	}
}

func TestReadMessage_FromStdin(t *testing.T) {
	stdin := strings.NewReader("piped message\n")
	msg, err := ReadMessage("", nil, stdin, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "piped message" {
		t.Errorf("got %q, want %q", msg, "piped message")
	}
}

func TestReadMessage_NoInput_Terminal(t *testing.T) {
	_, err := ReadMessage("", nil, nil, true)
	if err == nil {
		t.Fatal("expected error when no input and terminal")
	}
}

func TestReadMessage_EmptyStdin(t *testing.T) {
	stdin := strings.NewReader("")
	_, err := ReadMessage("", nil, stdin, false)
	if err == nil {
		t.Fatal("expected error for empty stdin")
	}
}

func TestReadMessage_FileNotFound(t *testing.T) {
	_, err := ReadMessage("/nonexistent/file.txt", nil, nil, true)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestReadMessage_Priority_FileOverArg(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "msg.txt")
	if err := os.WriteFile(path, []byte("from file"), 0644); err != nil {
		t.Fatal(err)
	}

	msg, err := ReadMessage(path, []string{"from arg"}, nil, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "from file" {
		t.Errorf("got %q, want %q (file should take priority)", msg, "from file")
	}
}

func TestReadMessage_MultipleArgs(t *testing.T) {
	msg, err := ReadMessage("", []string{"build", "finished", "successfully"}, nil, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "build finished successfully" {
		t.Errorf("got %q, want %q", msg, "build finished successfully")
	}
}

func TestReadMessage_Priority_ArgOverStdin(t *testing.T) {
	stdin := strings.NewReader("from stdin")
	msg, err := ReadMessage("", []string{"from arg"}, stdin, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "from arg" {
		t.Errorf("got %q, want %q (arg should take priority)", msg, "from arg")
	}
}
