package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config configuration structure
type Config struct {
	ServerURL     string
	AgentCardPath string
	APIPath       string
	Timeout       time.Duration
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		ServerURL:     "http://localhost:8080",
		AgentCardPath: "card",
		APIPath:       "/api",
		Timeout:       30 * time.Second,
	}
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv() *Config {
	config := DefaultConfig()

	// Load configuration from environment variables
	if serverURL := os.Getenv("A2A_SERVER_URL"); serverURL != "" {
		config.ServerURL = serverURL
	}

	if agentCardPath := os.Getenv("A2A_AGENT_CARD_PATH"); agentCardPath != "" {
		config.AgentCardPath = agentCardPath
	}

	if apiPath := os.Getenv("A2A_API_PATH"); apiPath != "" {
		config.APIPath = apiPath
	}

	if timeoutStr := os.Getenv("A2A_TIMEOUT_SECONDS"); timeoutStr != "" {
		if timeout, err := strconv.Atoi(timeoutStr); err == nil {
			config.Timeout = time.Duration(timeout) * time.Second
		}
	}

	return config
}

// String returns string representation of configuration
func (c *Config) String() string {
	return fmt.Sprintf("ServerURL: %s, AgentCardPath: %s, APIPath: %s, Timeout: %v",
		c.ServerURL, c.AgentCardPath, c.APIPath, c.Timeout)
} 