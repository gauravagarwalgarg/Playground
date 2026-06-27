package main

import (
	"fmt"
	"time"
)

// Builder Pattern: ServerConfig builder with method chaining.
// Useful when constructing complex objects with many optional parameters.

// ServerConfig is the product being built.
type ServerConfig struct {
	Host         string
	Port         int
	TLS          bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxConns     int
	LogLevel     string
}

func (c ServerConfig) String() string {
	return fmt.Sprintf(
		"Server{host=%s, port=%d, tls=%v, readTimeout=%v, writeTimeout=%v, maxConns=%d, logLevel=%s}",
		c.Host, c.Port, c.TLS, c.ReadTimeout, c.WriteTimeout, c.MaxConns, c.LogLevel,
	)
}

// ServerBuilder constructs ServerConfig step by step.
type ServerBuilder struct {
	config ServerConfig
}

func NewServerBuilder(host string, port int) *ServerBuilder {
	return &ServerBuilder{
		config: ServerConfig{
			Host:         host,
			Port:         port,
			TLS:          false,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxConns:     100,
			LogLevel:     "info",
		},
	}
}

func (b *ServerBuilder) WithTLS(enabled bool) *ServerBuilder {
	b.config.TLS = enabled
	return b
}

func (b *ServerBuilder) WithReadTimeout(d time.Duration) *ServerBuilder {
	b.config.ReadTimeout = d
	return b
}

func (b *ServerBuilder) WithWriteTimeout(d time.Duration) *ServerBuilder {
	b.config.WriteTimeout = d
	return b
}

func (b *ServerBuilder) WithMaxConns(n int) *ServerBuilder {
	b.config.MaxConns = n
	return b
}

func (b *ServerBuilder) WithLogLevel(level string) *ServerBuilder {
	b.config.LogLevel = level
	return b
}

// Build returns the final ServerConfig. Validates required fields.
func (b *ServerBuilder) Build() (ServerConfig, error) {
	if b.config.Host == "" {
		return ServerConfig{}, fmt.Errorf("host is required")
	}
	if b.config.Port <= 0 || b.config.Port > 65535 {
		return ServerConfig{}, fmt.Errorf("invalid port: %d", b.config.Port)
	}
	return b.config, nil
}

func main() {
	// Build a production server config using method chaining
	server, err := NewServerBuilder("api.example.com", 443).
		WithTLS(true).
		WithReadTimeout(30 * time.Second).
		WithWriteTimeout(60 * time.Second).
		WithMaxConns(1000).
		WithLogLevel("warn").
		Build()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Production: %s\n", server)

	// Build a dev server with defaults
	dev, _ := NewServerBuilder("localhost", 8080).Build()
	fmt.Printf("Development: %s\n", dev)
}
