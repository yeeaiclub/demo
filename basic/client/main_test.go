package main

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	
	if config.ServerURL != "http://localhost:8080" {
		t.Errorf("Expected ServerURL to be 'http://localhost:8080', got '%s'", config.ServerURL)
	}
	
	if config.AgentCardPath != "card" {
		t.Errorf("Expected AgentCardPath to be 'card', got '%s'", config.AgentCardPath)
	}
	
	if config.APIPath != "/api" {
		t.Errorf("Expected APIPath to be '/api', got '%s'", config.APIPath)
	}
	
	if config.Timeout != 30*time.Second {
		t.Errorf("Expected Timeout to be 30s, got '%v'", config.Timeout)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Test environment variable loading
	config := LoadConfigFromEnv()
	
	// Verify default values
	if config.ServerURL != "http://localhost:8080" {
		t.Errorf("Expected default ServerURL to be 'http://localhost:8080', got '%s'", config.ServerURL)
	}
}

func TestConfigString(t *testing.T) {
	config := &Config{
		ServerURL:     "http://test:8080",
		AgentCardPath: "test-card",
		APIPath:       "/test-api",
		Timeout:       60 * time.Second,
	}
	
	str := config.String()
	expected := "ServerURL: http://test:8080, AgentCardPath: test-card, APIPath: /test-api, Timeout: 1m0s"
	
	if str != expected {
		t.Errorf("Expected config string to be '%s', got '%s'", expected, str)
	}
}

func TestNewA2AClient(t *testing.T) {
	config := DefaultConfig()
	client := NewA2AClient(config)
	
	if client == nil {
		t.Error("Expected client creation to succeed, but got nil")
	}
	
	if client.config != config {
		t.Error("Expected client config to match input config")
	}
	
	if client.httpClient == nil {
		t.Error("Expected HTTP client to not be nil")
	}
	
	if client.resolver == nil {
		t.Error("Expected resolver to not be nil")
	}
	
	if client.client == nil {
		t.Error("Expected A2A client to not be nil")
	}
} 