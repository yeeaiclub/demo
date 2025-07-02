package main

import (
	"context"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
)

type QueueManager struct {
}

func NewQueueManager() *QueueManager {
	return &QueueManager{}
}
func (q QueueManager) Add(ctx context.Context, taskId string, queue *event.Queue) error {
	panic("implement me")
}

func (q QueueManager) Get(ctx context.Context, taskId string) (*event.Queue, error) {
	panic("implement me")
}

func (q QueueManager) Tap(ctx context.Context, taskId string) (*event.Queue, error) {
	panic("implement me")
}

func (q QueueManager) Close(ctx context.Context, taskId string) error {
	panic("implement me")
}

func (q QueueManager) CreateOrTap(ctx context.Context, taskId string) (*event.Queue, error) {
	return event.NewQueue(10), nil
}
