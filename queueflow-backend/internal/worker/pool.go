package worker

import (
	"log"
	"sync"

	"queueflow/internal/queue"
)

type WorkerPool struct {
	queue   queue.Queue
	workers int
	manager *Manager
}

func NewWorkerPool(q queue.Queue, workers int, manager *Manager) *WorkerPool {
	return &WorkerPool{
		queue:   q,
		workers: workers,
		manager: manager,
	}
}

func (wp *WorkerPool) Start() {
	var wg sync.WaitGroup

	log.Printf("Starting %d workers...", wp.workers)

	for i := 1; i <= wp.workers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			w := NewWorker(wp.queue, wp.manager)

			log.Printf("[Worker %d] Started", id)

			for {
				w.Process()
			}
		}(i)
	}

	wg.Wait()
}
