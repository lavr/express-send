//go:build newrelic

package apm

import (
	"net/http"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type nrProvider struct {
	app *newrelic.Application
}

// New creates a New Relic APM provider.
// Configuration is read from standard New Relic environment variables:
//   - NEW_RELIC_APP_NAME (required)
//   - NEW_RELIC_LICENSE_KEY (required)
//   - NEW_RELIC_ENABLED (true/false)
func New() Provider {
	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
	)
	if err != nil {
		return noopFallback{}
	}
	return &nrProvider{app: app}
}

func (p *nrProvider) WrapHandler(pattern string, h http.Handler) http.Handler {
	_, wrapped := newrelic.WrapHandle(p.app, pattern, h)
	return wrapped
}

func (p *nrProvider) Shutdown() {
	p.app.Shutdown(5 * time.Second)
}

type noopFallback struct{}

func (noopFallback) WrapHandler(_ string, h http.Handler) http.Handler { return h }
func (noopFallback) Shutdown()                                          {}
