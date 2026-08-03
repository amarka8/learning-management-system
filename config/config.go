// Package config defines configuration dependencies used by the application.
package config

// Config represents the configuration available to application components.
// The server does not require any configuration values yet, but keeping the
// dependency behind this interface makes later configuration additions explicit.
type Config interface {
	Port() string
	URL() string
}

// NullConfig is the no-op configuration used when no configuration values are
// required, including in tests.
type NullConfig struct{}

func (NullConfig) Port() string { return "" }

func (NullConfig) URL() string { return "" }
