package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestServer creates a Server with stub send/chat functions for testing.
func newTestServer(keys []ResolvedKey, opts ...func(*Config)) *Server {
	cfg := Config{
		Listen:   ":0",
		BasePath: "/api/v1",
		Keys:     keys,
	}
	for _, o := range opts {
		o(&cfg)
	}
	sendFn := func(ctx context.Context, p *SendPayload) (string, error) {
		return "test-sync-id", nil
	}
	chatResolver := func(chatID string) (string, error) {
		if chatID == "unknown-alias" {
			return "", fmt.Errorf("unknown chat alias %q", chatID)
		}
		return chatID, nil
	}
	return New(cfg, sendFn, chatResolver)
}

func doRequest(srv *Server, method, path string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	srv.srv.Handler.ServeHTTP(w, req)
	return w
}

func parseResponse(t *testing.T, w *httptest.ResponseRecorder) sendResponse {
	t.Helper()
	var resp sendResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v (body: %s)", err, w.Body.String())
	}
	return resp
}

// --- healthz ---

func TestHealthz(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "test", Key: "key1"}})
	w := doRequest(srv, "GET", "/healthz", nil, nil)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"ok":true`) {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

// --- auth ---

func TestAuth_BearerToken(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "myapp", Key: "secret123"}})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"Authorization": "Bearer secret123",
		"Content-Type":  "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAuth_XAPIKey(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "myapp", Key: "secret123"}})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "secret123",
		"Content-Type": "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAuth_NoCredentials(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "myapp", Key: "secret123"}})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"Content-Type": "application/json",
	})
	if w.Code != 401 {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

func TestAuth_WrongKey(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "myapp", Key: "secret123"}})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "wrong-key",
		"Content-Type": "application/json",
	})
	if w.Code != 403 {
		t.Fatalf("expected 403, got %d", w.Code)
	}
}

func TestAuth_BotSignature_Enabled(t *testing.T) {
	srv := newTestServer(nil, func(c *Config) {
		c.AllowBotSecretAuth = true
		c.BotSignature = "ABCDEF123456"
	})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-Bot-Signature": "ABCDEF123456",
		"Content-Type":    "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestAuth_BotSignature_Disabled(t *testing.T) {
	srv := newTestServer(nil, func(c *Config) {
		c.AllowBotSecretAuth = false
		c.BotSignature = "ABCDEF123456"
	})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-Bot-Signature": "ABCDEF123456",
		"Content-Type":    "application/json",
	})
	// 403 because credentials were presented but rejected (bot secret auth is disabled)
	if w.Code != 403 {
		t.Fatalf("expected 403 when bot secret auth disabled, got %d", w.Code)
	}
}

func TestAuth_BotSignature_Wrong(t *testing.T) {
	srv := newTestServer(nil, func(c *Config) {
		c.AllowBotSecretAuth = true
		c.BotSignature = "ABCDEF123456"
	})
	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-Bot-Signature": "WRONG",
		"Content-Type":    "application/json",
	})
	if w.Code != 403 {
		t.Fatalf("expected 403, got %d", w.Code)
	}
}

func TestAuth_MultipleKeys(t *testing.T) {
	srv := newTestServer([]ResolvedKey{
		{Name: "monitoring", Key: "mon-key"},
		{Name: "ci", Key: "ci-key"},
	})
	body := `{"chat_id":"chat-1","message":"hi"}`

	// Second key works
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "ci-key",
		"Content-Type": "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

// --- send handler: JSON ---

func TestSend_JSON_TextOnly(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"chat_id":"chat-1","message":"hello"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	resp := parseResponse(t, w)
	if !resp.OK {
		t.Fatalf("expected ok=true")
	}
	if resp.SyncID != "test-sync-id" {
		t.Fatalf("expected sync_id=test-sync-id, got %q", resp.SyncID)
	}
}

func TestSend_JSON_WithFile(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"chat_id":"chat-1","message":"see attached","file":{"name":"test.txt","data":"aGVsbG8="}}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestSend_JSON_MissingChatID(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"message":"hello"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	resp := parseResponse(t, w)
	if resp.Error != "chat_id is required" {
		t.Fatalf("unexpected error: %s", resp.Error)
	}
}

func TestSend_JSON_EmptyBody(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"chat_id":"chat-1"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
	resp := parseResponse(t, w)
	if resp.Error != "message or file required" {
		t.Fatalf("unexpected error: %s", resp.Error)
	}
}

func TestSend_JSON_InvalidStatus(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"chat_id":"chat-1","message":"hi","status":"bad"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestSend_JSON_InvalidJSON(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader("{invalid"), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestSend_JSON_ChatAlias_NotFound(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	body := `{"chat_id":"unknown-alias","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestSend_UnsupportedContentType(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader("data"), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "text/plain",
	})
	if w.Code != 415 {
		t.Fatalf("expected 415, got %d", w.Code)
	}
}

// --- send handler: multipart ---

func TestSend_Multipart_TextOnly(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("chat_id", "chat-1")
	w.WriteField("message", "multipart text")
	w.Close()

	rec := doRequest(srv, "POST", "/api/v1/send", &buf, map[string]string{
		"X-API-Key":    "k",
		"Content-Type": w.FormDataContentType(),
	})
	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestSend_Multipart_WithFile(t *testing.T) {
	srv := newTestServer([]ResolvedKey{{Name: "t", Key: "k"}})

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("chat_id", "chat-1")
	w.WriteField("message", "file attached")
	part, _ := w.CreateFormFile("file", "report.txt")
	part.Write([]byte("file content here"))
	w.Close()

	rec := doRequest(srv, "POST", "/api/v1/send", &buf, map[string]string{
		"X-API-Key":    "k",
		"Content-Type": w.FormDataContentType(),
	})
	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

// --- send handler: upstream error ---

func TestSend_UpstreamError(t *testing.T) {
	cfg := Config{
		Listen:   ":0",
		BasePath: "/api/v1",
		Keys:     []ResolvedKey{{Name: "t", Key: "k"}},
	}
	failSend := func(ctx context.Context, p *SendPayload) (string, error) {
		return "", fmt.Errorf("connection refused")
	}
	chatResolver := func(chatID string) (string, error) { return chatID, nil }
	srv := New(cfg, failSend, chatResolver)

	body := `{"chat_id":"chat-1","message":"hi"}`
	w := doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 502 {
		t.Fatalf("expected 502, got %d", w.Code)
	}
	resp := parseResponse(t, w)
	if !strings.Contains(resp.Error, "upstream error") {
		t.Fatalf("expected upstream error, got: %s", resp.Error)
	}
}

// --- base path ---

func TestSend_CustomBasePath(t *testing.T) {
	cfg := Config{
		Listen:   ":0",
		BasePath: "/express",
		Keys:     []ResolvedKey{{Name: "t", Key: "k"}},
	}
	sendFn := func(ctx context.Context, p *SendPayload) (string, error) { return "id", nil }
	chatResolver := func(chatID string) (string, error) { return chatID, nil }
	srv := New(cfg, sendFn, chatResolver)

	body := `{"chat_id":"chat-1","message":"hi"}`

	// Custom path works
	w := doRequest(srv, "POST", "/express/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code != 200 {
		t.Fatalf("expected 200 on /express/send, got %d: %s", w.Code, w.Body.String())
	}

	// Default path does not work
	w = doRequest(srv, "POST", "/api/v1/send", strings.NewReader(body), map[string]string{
		"X-API-Key":    "k",
		"Content-Type": "application/json",
	})
	if w.Code == 200 {
		t.Fatalf("expected /api/v1/send to NOT match with custom base path")
	}
}
