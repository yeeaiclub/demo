package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yeeaiclub/a2a-go/sdk/client"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

// A2AClient A2A client wrapper
type A2AClient struct {
	config     *Config
	httpClient *http.Client
	resolver   *client.A2ACardResolver
	client     *client.A2AClient
}

// NewA2AClient creates a new A2A client
func NewA2AClient(config *Config) *A2AClient {
	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	resolver := client.NewA2ACardResolver(
		httpClient, 
		config.ServerURL, 
		client.WithAgentCardPath(config.AgentCardPath),
	)

	a2aClient := client.NewClient(httpClient, config.ServerURL+config.APIPath)

	return &A2AClient{
		config:     config,
		httpClient: httpClient,
		resolver:   resolver,
		client:     a2aClient,
	}
}

// GetAgentCard gets agent card information
func (c *A2AClient) GetAgentCard() (*types.AgentCard, error) {
	log.Println("Getting agent card information...")
	card, err := c.resolver.GetAgentCard()
	if err != nil {
		return nil, fmt.Errorf("failed to get agent card: %w", err)
	}
	log.Printf("Successfully got agent card: %s (version: %s)", card.Name, card.Version)
	return &card, nil
}

// SendMessage sends a message
func (c *A2AClient) SendMessage(taskID, message string) (*types.Task, error) {
	log.Printf("Sending message to task %s: %s", taskID, message)
	
	resp, err := c.client.SendMessage(types.MessageSendParam{
		Message: &types.Message{
			TaskID: taskID,
			Role:   types.User,
			Parts: []types.Part{
				&types.TextPart{Kind: "text", Text: message},
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("server returned error: %v", resp.Error)
	}

	task, err := types.MapTo[types.Task](resp.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse task response: %w", err)
	}

	log.Printf("Message sent successfully, task ID: %s", task.Id)
	return &task, nil
}

// SendInteractiveMessage sends an interactive message
func (c *A2AClient) SendInteractiveMessage() error {
	fmt.Print("Enter message to send: ")
	var message string
	fmt.Scanln(&message)

	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	task, err := c.SendMessage("1", message)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Message sent, task ID: %s\n", task.Id)
	return nil
}

// RunDemo runs the demonstration
func (c *A2AClient) RunDemo() error {
	log.Println("🚀 Starting A2A client demonstration...")

	// 1. Get agent card
	card, err := c.GetAgentCard()
	if err != nil {
		return fmt.Errorf("demonstration failed: %w", err)
	}

	fmt.Printf("📋 Agent Information:\n")
	fmt.Printf("   Name: %s\n", card.Name)
	fmt.Printf("   Description: %s\n", card.Description)
	fmt.Printf("   Version: %s\n", card.Version)
	fmt.Println()

	// 2. Send default message
	log.Println("📤 Sending default message...")
	task, err := c.SendMessage("1", "hello, world")
	if err != nil {
		return fmt.Errorf("failed to send default message: %w", err)
	}
	fmt.Printf("✅ Default message sent successfully, task ID: %s\n", task.Id)
	fmt.Println()

	// 3. Interactive message sending
	fmt.Println("💬 Now you can send custom messages:")
	for {
		err := c.SendInteractiveMessage()
		if err != nil {
			log.Printf("❌ Failed to send message: %v", err)
			fmt.Print("Continue? (y/n): ")
			var choice string
			fmt.Scanln(&choice)
			if choice != "y" && choice != "Y" {
				break
			}
		}
	}

	return nil
}

func main() {
	// Set log format
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("🎯 A2A Go SDK Client Demo")

	// Create configuration
	config := LoadConfigFromEnv()

	// Create client
	client := NewA2AClient(config)

	// Run demonstration
	if err := client.RunDemo(); err != nil {
		log.Printf("❌ Demo failed: %v", err)
		os.Exit(1)
	}

	log.Println("🎉 Demo completed!")
}
