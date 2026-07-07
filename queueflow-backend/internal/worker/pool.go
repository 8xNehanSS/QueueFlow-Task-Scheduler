package worker

import (
	"database/sql"
	"log"
	"sync"

	"queueflow/internal/queue"
	"queueflow/internal/repository"
)

type WorkerPool struct {
	queue   queue.Queue
	workers int
	manager *Manager
	db      *sql.DB
}

func NewWorkerPool(q queue.Queue, workers int, manager *Manager, db *sql.DB) *WorkerPool {
	return &WorkerPool{
		queue:   q,
		workers: workers,
		manager: manager,
		db:      db,
	}
}

func (wp *WorkerPool) Start() {
	var wg sync.WaitGroup

	log.Printf("Starting %d workers...", wp.workers)

	for i := 1; i <= wp.workers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			repo := repository.NewJobRepository(wp.db)
			w := NewWorker(wp.queue, wp.manager, repo)

			log.Printf("[Worker %d] Started", id)

			for {
				w.Process()
			}
		}(i)
	}

	wg.Wait()
}
