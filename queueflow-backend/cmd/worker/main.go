package main

import (
	"log"
	"queueflow/internal/queue"
	"queueflow/internal/worker"
)

func main() {
	q := queue.NewRedisQueue()
	w := worker.NewWorker(q)

	log.Println("Worker started...")

	for {
		w.Process()
	}
}
