package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_FromYAML(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  main:
    host: express.example.com
    id: bot-123
    secret: my-secret
chats:
  deploy: chat-456
cache:
  type: file
  file_path: /tmp/token.cache
  ttl: 1800
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(Flags{ConfigPath: cfgPath})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Host != "express.example.com" {
		t.Errorf("Host = %q, want %q", cfg.Host, "express.example.com")
	}
	if cfg.BotID != "bot-123" {
		t.Errorf("BotID = %q, want %q", cfg.BotID, "bot-123")
	}
	if cfg.BotSecret != "my-secret" {
		t.Errorf("BotSecret = %q, want %q", cfg.BotSecret, "my-secret")
	}
	if cfg.BotName != "main" {
		t.Errorf("BotName = %q, want %q", cfg.BotName, "main")
	}
	if cfg.Cache.Type != "file" {
		t.Errorf("Cache.Type = %q, want %q", cfg.Cache.Type, "file")
	}
	if cfg.Cache.TTL != 1800 {
		t.Errorf("Cache.TTL = %d, want %d", cfg.Cache.TTL, 1800)
	}
}

func TestLoad_EnvOverridesBot(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  main:
    host: from-yaml.com
    id: yaml-bot
    secret: yaml-secret
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("EXPRESS_HOST", "from-env.com")
	t.Setenv("EXPRESS_BOT_ID", "env-bot")

	cfg, err := Load(Flags{ConfigPath: cfgPath})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Host != "from-env.com" {
		t.Errorf("Host = %q, want %q (env should override bot)", cfg.Host, "from-env.com")
	}
	if cfg.BotID != "env-bot" {
		t.Errorf("BotID = %q, want %q (env should override bot)", cfg.BotID, "env-bot")
	}
	if cfg.BotSecret != "yaml-secret" {
		t.Errorf("BotSecret = %q, want %q", cfg.BotSecret, "yaml-secret")
	}
}

func TestLoad_FlagsOverrideAll(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  main:
    host: from-yaml.com
    id: yaml-bot
    secret: yaml-secret
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("EXPRESS_HOST", "from-env.com")

	cfg, err := Load(Flags{
		ConfigPath: cfgPath,
		Host:       "from-flag.com",
	})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Host != "from-flag.com" {
		t.Errorf("Host = %q, want %q (flag should override all)", cfg.Host, "from-flag.com")
	}
}

func TestLoad_MultipleBots_RequiresFlag(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  prod:
    host: prod.com
    id: prod-bot
    secret: prod-secret
  test:
    host: test.com
    id: test-bot
    secret: test-secret
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(Flags{ConfigPath: cfgPath})
	if err == nil {
		t.Fatal("expected error for multiple bots without --bot")
	}
	if !strings.Contains(err.Error(), "multiple bots") {
		t.Errorf("error = %q, should mention multiple bots", err.Error())
	}
}

func TestLoad_MultipleBots_WithFlag(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  prod:
    host: prod.com
    id: prod-bot
    secret: prod-secret
  test:
    host: test.com
    id: test-bot
    secret: test-secret
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(Flags{ConfigPath: cfgPath, Bot: "prod"})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if cfg.Host != "prod.com" {
		t.Errorf("Host = %q, want %q", cfg.Host, "prod.com")
	}
	if cfg.BotID != "prod-bot" {
		t.Errorf("BotID = %q, want %q", cfg.BotID, "prod-bot")
	}
	if cfg.BotName != "prod" {
		t.Errorf("BotName = %q, want %q", cfg.BotName, "prod")
	}
}

func TestLoad_UnknownBot(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  prod:
    host: prod.com
    id: prod-bot
    secret: prod-secret
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(Flags{ConfigPath: cfgPath, Bot: "staging"})
	if err == nil {
		t.Fatal("expected error for unknown bot")
	}
	if !strings.Contains(err.Error(), "staging") {
		t.Errorf("error = %q, should mention staging", err.Error())
	}
}

func TestLoad_MissingRequired(t *testing.T) {
	tests := []struct {
		name  string
		flags Flags
		want  string
	}{
		{"no host", Flags{BotID: "b", Secret: "s"}, "host is required"},
		{"no bot_id", Flags{Host: "h", Secret: "s"}, "bot id is required"},
		{"no secret", Flags{Host: "h", BotID: "b"}, "bot secret is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(tt.flags)
			if err == nil {
				t.Fatal("expected error")
			}
			if got := err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("error = %q, should contain %q", got, tt.want)
			}
		})
	}
}

func TestCacheKey(t *testing.T) {
	cfg := &Config{Host: "express.example.com", BotID: "bot-123"}
	want := "express.example.com:bot-123"
	if got := cfg.CacheKey(); got != want {
		t.Errorf("CacheKey() = %q, want %q", got, want)
	}
}

func TestRequireChatID(t *testing.T) {
	cfg := &Config{ChatID: ""}
	if err := cfg.RequireChatID(); err == nil {
		t.Fatal("expected error for empty ChatID")
	}

	cfg.ChatID = "some-chat"
	cfg.Chats = map[string]string{"some-chat": "uuid-123"}
	if err := cfg.RequireChatID(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ChatID != "uuid-123" {
		t.Errorf("ChatID = %q, want resolved %q", cfg.ChatID, "uuid-123")
	}
}

func TestRequireChatID_SingleAlias(t *testing.T) {
	cfg := &Config{Chats: map[string]string{"deploy": "uuid-456"}}
	if err := cfg.RequireChatID(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.ChatID != "uuid-456" {
		t.Errorf("ChatID = %q, want %q", cfg.ChatID, "uuid-456")
	}
}

func TestLoad_NoCacheFlag(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	content := `
bots:
  main:
    host: h
    id: b
    secret: s
cache:
  type: file
`
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(Flags{ConfigPath: cfgPath, NoCache: true})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Cache.Type != "none" {
		t.Errorf("Cache.Type = %q, want %q (--no-cache should disable)", cfg.Cache.Type, "none")
	}
}

func TestLoad_MissingConfigFile_NotExplicit(t *testing.T) {
	cfg, err := Load(Flags{
		Host:   "h",
		BotID:  "b",
		Secret: "s",
	})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Host != "h" {
		t.Errorf("Host = %q, want %q", cfg.Host, "h")
	}
}

func TestLoad_MissingConfigFile_Explicit(t *testing.T) {
	_, err := Load(Flags{ConfigPath: "/nonexistent/config.yaml"})
	if err == nil {
		t.Fatal("expected error for explicitly specified missing config")
	}
}

func TestLoad_DefaultCacheTTL(t *testing.T) {
	cfg, err := Load(Flags{
		Host:   "h",
		BotID:  "b",
		Secret: "s",
	})
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Cache.TTL != 3600 {
		t.Errorf("Cache.TTL = %d, want 3600 (default)", cfg.Cache.TTL)
	}
}
