package apm

import "net/http"

// Provider is an APM provider. The implementation is selected at compile time via build tags.
type Provider interface {
	// WrapHandler wraps an HTTP handler for tracing.
	WrapHandler(pattern string, h http.Handler) http.Handler
	// Shutdown flushes pending data and stops the agent.
	Shutdown()
}
