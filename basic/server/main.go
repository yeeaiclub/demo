package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yeeaiclub/a2a-go/sdk/server/handler"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

// Server A2A server wrapper
type Server struct {
	config *Config
	server *handler.Server
}

// NewServer creates a new A2A server
func NewServer(config *Config) *Server {
	// Create agent card
	card := types.AgentCard{
		Name:        config.AgentName,
		Description: config.AgentDesc,
		Version:     config.AgentVersion,
	}

	// Create task store
	mem := tasks.NewInMemoryTaskStore()
	
	// Pre-create some example tasks
	ctx := context.Background()
	mem.Save(ctx, &types.Task{Id: "1"})
	mem.Save(ctx, &types.Task{Id: "demo-task"})

	// Create executor
	executor := NewPrintExecutor(mem)
	
	// Create queue manager
	queue := NewQueueManager()
	
	// Create default handler
	defaultHandler := handler.NewDefaultHandler(
		mem, 
		executor, 
		handler.WithQueueManger(queue),
	)

	// Create server
	server := handler.NewServer(config.CardPath, config.APIPath, card, defaultHandler)

	return &Server{
		config: config,
		server: server,
	}
}

// Start starts the server
func (s *Server) Start() error {
	log.Printf("🚀 Starting A2A server...")
	log.Printf("📋 Agent Information:")
	log.Printf("   Name: %s", s.config.AgentName)
	log.Printf("   Description: %s", s.config.AgentDesc)
	log.Printf("   Version: %s", s.config.AgentVersion)
	log.Printf("🌐 Server URL: http://localhost:%d", s.config.Port)
	log.Printf("📄 Agent Card: http://localhost:%d%s", s.config.Port, s.config.CardPath)
	log.Printf("🔌 API Endpoint: http://localhost:%d%s", s.config.Port, s.config.APIPath)
	log.Println()

	// Start server
	go func() {
		if err := s.server.Start(s.config.Port); err != nil {
			log.Printf("❌ Server failed to start: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("\n🛑 Received interrupt signal, shutting down server...")
	return nil
}

func main() {
	// Set log format
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("🎯 A2A Go SDK Server Demo")

	// Create configuration
	config := LoadConfigFromEnv()

	// Create server
	server := NewServer(config)

	// Start server
	if err := server.Start(); err != nil {
		log.Printf("❌ Server failed: %v", err)
		os.Exit(1)
	}

	log.Println("🎉 Server shutdown complete!")
}
