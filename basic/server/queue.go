package main

import (
	"context"
	"log"
	"sync"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
)

// QueueManager queue manager
type QueueManager struct {
	queues map[string]*event.Queue
	mutex  sync.RWMutex
}

// NewQueueManager creates a new queue manager
func NewQueueManager() *QueueManager {
	return &QueueManager{
		queues: make(map[string]*event.Queue),
	}
}

// Add adds a queue
func (q *QueueManager) Add(ctx context.Context, taskId string, queue *event.Queue) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	
	q.queues[taskId] = queue
	log.Printf("➕ Added queue: %s", taskId)
	return nil
}

// Get gets a queue
func (q *QueueManager) Get(ctx context.Context, taskId string) (*event.Queue, error) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	
	queue, exists := q.queues[taskId]
	if !exists {
		log.Printf("⚠️  Queue not found: %s", taskId)
		return nil, nil
	}
	
	log.Printf("📋 Retrieved queue: %s", taskId)
	return queue, nil
}

// Tap taps a queue (get without removing)
func (q *QueueManager) Tap(ctx context.Context, taskId string) (*event.Queue, error) {
	return q.Get(ctx, taskId)
}

// Close closes a queue
func (q *QueueManager) Close(ctx context.Context, taskId string) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	
	if _, exists := q.queues[taskId]; exists {
		delete(q.queues, taskId)
		log.Printf("🔒 Closed queue: %s", taskId)
	}
	return nil
}

// CreateOrTap creates or gets a queue
func (q *QueueManager) CreateOrTap(ctx context.Context, taskId string) (*event.Queue, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	
	// Check if queue already exists
	if queue, exists := q.queues[taskId]; exists {
		log.Printf("📋 Retrieved existing queue: %s", taskId)
		return queue, nil
	}
	
	// Create new queue
	queue := event.NewQueue(10)
	q.queues[taskId] = queue
	log.Printf("🆕 Created new queue: %s", taskId)
	return queue, nil
}
