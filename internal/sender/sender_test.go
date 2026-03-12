package sender

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSend(t *testing.T) {
	chatID := "054af49e-5e18-4dca-ad73-4f96b6de63fa"
	token := "test-token"
	message := "Hello, world!"

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/api/v4/botx/notifications/direct" {
			t.Errorf("path = %q, want /api/v4/botx/notifications/direct", r.URL.Path)
		}
		if auth := r.Header.Get("Authorization"); auth != "Bearer "+token {
			t.Errorf("Authorization = %q, want %q", auth, "Bearer "+token)
		}

		var payload notificationRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decoding request body: %v", err)
		}
		if payload.GroupChatID != chatID {
			t.Errorf("group_chat_id = %q, want %q", payload.GroupChatID, chatID)
		}
		if payload.Notification.Body != message {
			t.Errorf("body = %q, want %q", payload.Notification.Body, message)
		}
		if payload.Notification.Status != "ok" {
			t.Errorf("status = %q, want %q", payload.Notification.Status, "ok")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sendResponse{Status: "ok"})
	}))
	defer srv.Close()

	ctx := context.Background()
	err := sendWithClient(ctx, srv.URL, token, chatID, message, srv.Client())
	if err != nil {
		t.Fatalf("Send() error: %v", err)
	}
}

func TestSend_Unauthorized(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	ctx := context.Background()
	err := sendWithClient(ctx, srv.URL, "bad-token", "chat-id", "msg", srv.Client())
	if !errors.Is(err, ErrUnauthorized) {
		t.Errorf("expected ErrUnauthorized, got: %v", err)
	}
}

func TestSend_ServerError(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	ctx := context.Background()
	err := sendWithClient(ctx, srv.URL, "token", "chat-id", "msg", srv.Client())
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}
