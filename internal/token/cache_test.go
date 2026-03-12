package token

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
)

func TestNoopCache(t *testing.T) {
	ctx := context.Background()
	c := NoopCache{}

	if err := c.Set(ctx, "key", "token", time.Hour); err != nil {
		t.Fatalf("Set() error: %v", err)
	}

	val, err := c.Get(ctx, "key")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "" {
		t.Errorf("Get() = %q, want empty (noop)", val)
	}
}

func TestFileCache(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tokens.json")
	c := &FileCache{Path: path}
	ctx := context.Background()

	// Miss on empty cache
	val, err := c.Get(ctx, "bot1")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "" {
		t.Errorf("Get() = %q, want empty on miss", val)
	}

	// Set and get
	if err := c.Set(ctx, "bot1", "token-abc", time.Hour); err != nil {
		t.Fatalf("Set() error: %v", err)
	}

	val, err = c.Get(ctx, "bot1")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "token-abc" {
		t.Errorf("Get() = %q, want %q", val, "token-abc")
	}

	// Different key is still a miss
	val, err = c.Get(ctx, "bot2")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "" {
		t.Errorf("Get(bot2) = %q, want empty", val)
	}
}

func TestFileCache_Expiry(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tokens.json")
	c := &FileCache{Path: path}
	ctx := context.Background()

	// Set with very short TTL
	if err := c.Set(ctx, "bot1", "token-abc", time.Millisecond); err != nil {
		t.Fatalf("Set() error: %v", err)
	}

	time.Sleep(5 * time.Millisecond)

	val, err := c.Get(ctx, "bot1")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "" {
		t.Errorf("Get() = %q, want empty (expired)", val)
	}
}

func TestFileCache_MultipleKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tokens.json")
	c := &FileCache{Path: path}
	ctx := context.Background()

	c.Set(ctx, "bot1", "token1", time.Hour)
	c.Set(ctx, "bot2", "token2", time.Hour)

	val1, _ := c.Get(ctx, "bot1")
	val2, _ := c.Get(ctx, "bot2")

	if val1 != "token1" {
		t.Errorf("bot1 = %q, want %q", val1, "token1")
	}
	if val2 != "token2" {
		t.Errorf("bot2 = %q, want %q", val2, "token2")
	}
}

func TestVaultCache(t *testing.T) {
	storedData := make(map[string]interface{})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Vault-Token") != "test-token" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		switch r.Method {
		case http.MethodGet:
			if len(storedData) == 0 {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"data": map[string]interface{}{
					"data": storedData,
				},
			})
		case http.MethodPost:
			var body struct {
				Data map[string]interface{} `json:"data"`
			}
			json.NewDecoder(r.Body).Decode(&body)
			storedData = body.Data
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer srv.Close()

	c := &VaultCache{
		URL:   srv.URL,
		Path:  "secret/data/tokens",
		Token: "test-token",
	}
	ctx := context.Background()

	// Miss on empty
	val, err := c.Get(ctx, "bot1")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "" {
		t.Errorf("Get() = %q, want empty on miss", val)
	}

	// Set and get
	if err := c.Set(ctx, "bot1", "vault-token-123", time.Hour); err != nil {
		t.Fatalf("Set() error: %v", err)
	}

	val, err = c.Get(ctx, "bot1")
	if err != nil {
		t.Fatalf("Get() error: %v", err)
	}
	if val != "vault-token-123" {
		t.Errorf("Get() = %q, want %q", val, "vault-token-123")
	}
}
