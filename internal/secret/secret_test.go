package secret

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResolve_Literal(t *testing.T) {
	val, err := Resolve("my-secret-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "my-secret-key" {
		t.Errorf("got %q, want %q", val, "my-secret-key")
	}
}

func TestResolve_Env(t *testing.T) {
	t.Setenv("TEST_SECRET_VAR", "env-secret-value")

	val, err := Resolve("env:TEST_SECRET_VAR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "env-secret-value" {
		t.Errorf("got %q, want %q", val, "env-secret-value")
	}
}

func TestResolve_Env_Empty(t *testing.T) {
	t.Setenv("TEST_EMPTY_VAR", "")

	_, err := Resolve("env:TEST_EMPTY_VAR")
	if err == nil {
		t.Fatal("expected error for empty env var")
	}
}

func TestResolve_Env_Unset(t *testing.T) {
	_, err := Resolve("env:DEFINITELY_NOT_SET_VAR_12345")
	if err == nil {
		t.Fatal("expected error for unset env var")
	}
}

func TestResolve_Vault(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/secret/data/myapp" {
			t.Errorf("path = %q, want /v1/secret/data/myapp", r.URL.Path)
		}
		if tok := r.Header.Get("X-Vault-Token"); tok != "test-vault-token" {
			t.Errorf("X-Vault-Token = %q, want %q", tok, "test-vault-token")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"data": map[string]interface{}{
					"bot_secret": "vault-secret-value",
				},
			},
		})
	}))
	defer srv.Close()

	t.Setenv("VAULT_ADDR", srv.URL)
	t.Setenv("VAULT_TOKEN", "test-vault-token")

	val, err := Resolve("vault:secret/data/myapp#bot_secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "vault-secret-value" {
		t.Errorf("got %q, want %q", val, "vault-secret-value")
	}
}

func TestResolve_Vault_BadSpec(t *testing.T) {
	_, err := Resolve("vault:no-hash-sign")
	if err == nil {
		t.Fatal("expected error for spec without #")
	}
}

func TestResolve_Vault_NoAddr(t *testing.T) {
	t.Setenv("VAULT_ADDR", "")
	t.Setenv("VAULT_TOKEN", "tok")

	_, err := Resolve("vault:path#key")
	if err == nil {
		t.Fatal("expected error when VAULT_ADDR not set")
	}
}

func TestResolve_Vault_KeyNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"data": map[string]interface{}{
					"other_key": "value",
				},
			},
		})
	}))
	defer srv.Close()

	t.Setenv("VAULT_ADDR", srv.URL)
	t.Setenv("VAULT_TOKEN", "tok")

	_, err := Resolve("vault:secret/data/app#missing_key")
	if err == nil {
		t.Fatal("expected error for missing key")
	}
}
