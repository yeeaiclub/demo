package main

import (
	"context"
	"fmt"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
	"github.com/yeeaiclub/a2a-go/sdk/server/execution"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks/updater"
)

type PrintExecutor struct {
	store tasks.TaskStore
}

func NewPrintExecutor(store tasks.TaskStore) *PrintExecutor {
	return &PrintExecutor{store: store}
}

func (m PrintExecutor) Execute(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	fmt.Println("hello, word")
	fmt.Println(requestContext.Params.Message.Parts[0])
	fmt.Println(requestContext.TaskId)
	fmt.Println(requestContext.Params.Message.Role)
	taskUpdater := updater.NewTaskUpdater(queue, requestContext.TaskId, requestContext.ContextId)
	taskUpdater.Complete(updater.WithFinal(true))
	return nil
}

func (m PrintExecutor) Cancel(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	return nil
}
