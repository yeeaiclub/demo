package main

import (
	"context"
	"fmt"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
	"github.com/yeeaiclub/a2a-go/sdk/server/execution"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks"
	"github.com/yeeaiclub/a2a-go/sdk/server/tasks/updater"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

type PrintExecutor struct {
	store tasks.TaskStore
}

func NewPrintExecutor(store tasks.TaskStore) *PrintExecutor {
	return &PrintExecutor{store: store}
}

func (m PrintExecutor) Execute(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	fmt.Println("hello, word")

	task, err := m.store.Get(ctx, requestContext.TaskId)
	if err != nil {
		return err
	}

	if task == nil {
		err = m.store.Save(ctx, &types.Task{Id: requestContext.TaskId, ContextId: requestContext.ContextId})
		if err != nil {
			return err
		}
	}
	taskUpdater := updater.NewTaskUpdater(queue, requestContext.TaskId, requestContext.ContextId)
	taskUpdater.Complete(updater.WithFinal(true))
	return nil
}

func (m PrintExecutor) Cancel(ctx context.Context, requestContext *execution.RequestContext, queue *event.Queue) error {
	return nil
}
