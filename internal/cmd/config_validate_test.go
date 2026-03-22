package cmd

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/lavr/express-botx/internal/config"
)

func TestConfigValidate_ValidConfig(t *testing.T) {
	cfg := `bots:
  mybot:
    host: express.example.com
    id: 550e8400-e29b-41d4-a716-446655440000
    secret: my-secret
chats:
  mychat:
    id: 660e8400-e29b-41d4-a716-446655440000
    bot: mybot
`
	path := writeTestConfig(t, cfg)
	deps, stdout, _ := testDeps()

	err := runConfigValidate([]string{"--config", path}, deps)
	if err != nil {
		t.Fatalf("expected no error for valid config, got: %v", err)
	}
	if !strings.Contains(stdout.String(), "0 errors, 0 warnings") {
		t.Fatalf("expected 0 errors 0 warnings, got: %s", stdout.String())
	}
}

func TestConfigValidate_InvalidConfig_MissingRequired(t *testing.T) {
	cfg := `bots:
  mybot:
    host: express.example.com
`
	path := writeTestConfig(t, cfg)
	deps, stdout, _ := testDeps()

	err := runConfigValidate([]string{"--config", path}, deps)
	if err == nil {
		t.Fatal("expected error for invalid config")
	}
	if !strings.Contains(err.Error(), "validation failed") {
		t.Fatalf("expected validation failed error, got: %v", err)
	}
	output := stdout.String()
	if !strings.Contains(output, "[ERROR]") {
		t.Fatalf("expected ERROR in output, got: %s", output)
	}
	if !strings.Contains(output, "id is required") {
		t.Fatalf("expected 'id is required' in output, got: %s", output)
	}
}

func TestConfigValidate_WarningsOnly(t *testing.T) {
	cfg := `bots:
  mybot:
    host: express.example.com
    id: 550e8400-e29b-41d4-a716-446655440000
    secret: my-secret
unknown_top_key: something
`
	path := writeTestConfig(t, cfg)
	deps, stdout, _ := testDeps()

	err := runConfigValidate([]string{"--config", path}, deps)
	if err != nil {
		t.Fatalf("expected no error for warning-only config, got: %v", err)
	}
	output := stdout.String()
	if !strings.Contains(output, "[WARN]") {
		t.Fatalf("expected WARN in output, got: %s", output)
	}
	if !strings.Contains(output, "0 errors") {
		t.Fatalf("expected 0 errors in output, got: %s", output)
	}
}

func TestConfigValidate_JSONOutput(t *testing.T) {
	cfg := `bots:
  mybot:
    host: express.example.com
    id: not-a-uuid
    secret: my-secret
`
	path := writeTestConfig(t, cfg)
	deps, stdout, _ := testDeps()

	err := runConfigValidate([]string{"--config", path, "--format", "json"}, deps)
	if err == nil {
		t.Fatal("expected error for invalid config")
	}

	var results []config.ValidationResult
	if err := json.Unmarshal(stdout.Bytes(), &results); err != nil {
		t.Fatalf("expected valid JSON output, got error: %v, output: %s", err, stdout.String())
	}
	if len(results) == 0 {
		t.Fatal("expected at least one validation result")
	}
	found := false
	for _, r := range results {
		if r.Level == config.ValidationError && strings.Contains(r.Message, "UUID") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected UUID format error in results: %+v", results)
	}
}

func TestConfigValidate_MissingConfigFile(t *testing.T) {
	deps, _, _ := testDeps()

	err := runConfigValidate([]string{"--config", "/nonexistent/path/config.yaml"}, deps)
	if err == nil {
		t.Fatal("expected error for missing config file")
	}
}

func TestConfigValidate_SummaryLine(t *testing.T) {
	cfg := `bots:
  mybot:
    host: express.example.com
    id: not-a-uuid
    secret: my-secret
unknown_key: value
`
	path := writeTestConfig(t, cfg)
	deps, stdout, _ := testDeps()

	err := runConfigValidate([]string{"--config", path}, deps)
	if err == nil {
		t.Fatal("expected error for config with invalid UUID")
	}
	output := stdout.String()
	if !strings.Contains(output, "1 errors") {
		t.Fatalf("expected 1 errors in summary, got: %s", output)
	}
	if !strings.Contains(output, "1 warnings") {
		t.Fatalf("expected 1 warnings in summary, got: %s", output)
	}
}

func TestConfigValidate_Help(t *testing.T) {
	var stderr bytes.Buffer
	deps := Deps{
		Stdout:     &bytes.Buffer{},
		Stderr:     &stderr,
		Stdin:      strings.NewReader(""),
		IsTerminal: false,
	}

	err := runConfigValidate([]string{"--help"}, deps)
	if err != nil {
		t.Fatalf("expected no error for --help, got: %v", err)
	}
	if !strings.Contains(stderr.String(), "Validate config file") {
		t.Fatalf("expected usage text, got: %s", stderr.String())
	}
}
