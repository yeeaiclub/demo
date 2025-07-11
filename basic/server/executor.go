package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
	"github.com/yeeaiclub/a2a-go/sdk/server/execution"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks/updater"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

// PrintExecutor print executor
type PrintExecutor struct {
	store tasks.TaskStore
}

// NewPrintExecutor creates a new print executor
func NewPrintExecutor(store tasks.TaskStore) *PrintExecutor {
	return &PrintExecutor{store: store}
}

// Execute executes a task
func (m *PrintExecutor) Execute(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	log.Printf("🔄 Starting task execution: %s", requestContext.TaskId)
	log.Printf("📝 Message content: %s", requestContext.Params.Message.Parts[0])
	log.Printf("👤 User role: %s", requestContext.Params.Message.Role)

	// Print message content
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("📨 New message received:")
	
	// Extract and print text content
	for i, part := range requestContext.Params.Message.Parts {
		if textPart, ok := part.(*types.TextPart); ok {
			fmt.Printf("   Message %d: %s\n", i+1, textPart.Text)
		}
	}
	
	fmt.Printf("   Task ID: %s\n", requestContext.TaskId)
	fmt.Printf("   User Role: %s\n", requestContext.Params.Message.Role)
	fmt.Println(strings.Repeat("=", 50))

	// Create task updater
	taskUpdater := updater.NewTaskUpdater(queue, requestContext.TaskId, requestContext.ContextId)
	
	// Mark task as complete
	taskUpdater.Complete(updater.WithFinal(true))
	
	log.Printf("✅ Task execution completed: %s", requestContext.TaskId)
	return nil
}

// Cancel cancels a task
func (m *PrintExecutor) Cancel(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	log.Printf("❌ Cancelling task: %s", requestContext.TaskId)
	return nil
}
