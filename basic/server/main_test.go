package main

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	
	if config.Port != 8080 {
		t.Errorf("Expected Port to be 8080, got %d", config.Port)
	}
	
	if config.CardPath != "/card" {
		t.Errorf("Expected CardPath to be '/card', got '%s'", config.CardPath)
	}
	
	if config.APIPath != "/api" {
		t.Errorf("Expected APIPath to be '/api', got '%s'", config.APIPath)
	}
	
	if config.AgentName != "Print Agent" {
		t.Errorf("Expected AgentName to be 'Print Agent', got '%s'", config.AgentName)
	}
	
	if config.AgentVersion != "v0.1.0" {
		t.Errorf("Expected AgentVersion to be 'v0.1.0', got '%s'", config.AgentVersion)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Test environment variable loading
	config := LoadConfigFromEnv()
	
	// Verify default values
	if config.Port != 8080 {
		t.Errorf("Expected default Port to be 8080, got %d", config.Port)
	}
}

func TestConfigString(t *testing.T) {
	config := &Config{
		Port:          9090,
		CardPath:      "/test-card",
		APIPath:       "/test-api",
		AgentName:     "Test Agent",
		AgentVersion:  "v1.0.0",
	}
	
	str := config.String()
	expected := "Port: 9090, CardPath: /test-card, APIPath: /test-api, AgentName: Test Agent, AgentVersion: v1.0.0"
	
	if str != expected {
		t.Errorf("Expected config string to be '%s', got '%s'", expected, str)
	}
}

func TestNewServer(t *testing.T) {
	config := DefaultConfig()
	server := NewServer(config)
	
	if server == nil {
		t.Error("Expected server creation to succeed, but got nil")
	}
	
	if server.config != config {
		t.Error("Expected server config to match input config")
	}
	
	if server.server == nil {
		t.Error("Expected internal server to not be nil")
	}
}

func TestNewPrintExecutor(t *testing.T) {
	// This requires mocking TaskStore, skip for now
	// In a real project, you can use mock libraries to test
	t.Skip("Requires TaskStore mocking to test")
}

func TestNewQueueManager(t *testing.T) {
	queueManager := NewQueueManager()
	
	if queueManager == nil {
		t.Error("Expected queue manager creation to succeed, but got nil")
	}
	
	if queueManager.queues == nil {
		t.Error("Expected queues map to not be nil")
	}
} 