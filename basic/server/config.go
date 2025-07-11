package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config server configuration
type Config struct {
	Port          int
	CardPath      string
	APIPath       string
	AgentName     string
	AgentDesc     string
	AgentVersion  string
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Port:          8080,
		CardPath:      "/card",
		APIPath:       "/api",
		AgentName:     "Print Agent",
		AgentDesc:     "A simple print agent for A2A Go SDK demonstration",
		AgentVersion:  "v0.1.0",
	}
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv() *Config {
	config := DefaultConfig()

	// Load configuration from environment variables
	if portStr := os.Getenv("A2A_SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			config.Port = port
		}
	}

	if cardPath := os.Getenv("A2A_CARD_PATH"); cardPath != "" {
		config.CardPath = cardPath
	}

	if apiPath := os.Getenv("A2A_API_PATH"); apiPath != "" {
		config.APIPath = apiPath
	}

	if agentName := os.Getenv("A2A_AGENT_NAME"); agentName != "" {
		config.AgentName = agentName
	}

	if agentDesc := os.Getenv("A2A_AGENT_DESC"); agentDesc != "" {
		config.AgentDesc = agentDesc
	}

	if agentVersion := os.Getenv("A2A_AGENT_VERSION"); agentVersion != "" {
		config.AgentVersion = agentVersion
	}

	return config
}

// String returns string representation of configuration
func (c *Config) String() string {
	return fmt.Sprintf("Port: %d, CardPath: %s, APIPath: %s, AgentName: %s, AgentVersion: %s",
		c.Port, c.CardPath, c.APIPath, c.AgentName, c.AgentVersion)
} 