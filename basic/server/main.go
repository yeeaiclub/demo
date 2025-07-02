package main

import (
	"context"

	"github.com/yeeaiclub/a2a-go/sdk/server/handler"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

var card = types.AgentCard{
	Name:        "print card",
	Description: "print hello, word",
	Version:     "v0.1.0",
}

func main() {
	mem := tasks.NewInMemoryTaskStore()
	mem.Save(context.Background(), &types.Task{Id: "1"})
	executor := NewPrintExecutor(mem)
	queue := NewQueueManager()
	defaultHandler := handler.NewDefaultHandler(mem, executor, handler.WithQueueManger(queue))

	server := handler.NewServer("/card", "/api", card, defaultHandler)
	server.Start(8080)
}
