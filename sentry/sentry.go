package sentry

import (
	"github.com/getsentry/sentry-go"
)

type Component struct {
	DSN              string
	AttachStacktrace bool
	Environment      string
}

func NewComponent(dsn string, attachStacktrace bool, environment string) *Component {
	return &Component{
		DSN:              dsn,
		AttachStacktrace: attachStacktrace,
		Environment:      environment,
	}
}

// Start implements the Component interface.
func (c *Component) Start() error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              c.DSN,
		AttachStacktrace: c.AttachStacktrace,
		Environment:      c.Environment,
	})
}

// Stop implements the Component interface.
func (c *Component) Stop() error {
	sentry.Flush(2)

	return nil
}

// Health implements the Component interface.
func (c *Component) Health() error {
	return nil
}

// Name implements the Component interface.
func (c *Component) Name() string {
	return "sentry"
}
