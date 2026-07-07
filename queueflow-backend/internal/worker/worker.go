package worker

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/queue"
)

type Worker struct {
	queue   queue.Queue
	manager *Manager
}

func NewWorker(q queue.Queue, m *Manager) *Worker {
	return &Worker{
		queue:   q,
		manager: m,
	}
}

func (w *Worker) Process() {

	job, err := w.queue.Pop()

	if err != nil {
		log.Println("queue error:", err)
		return
	}

	jobType := job.Type

	handler, exists := w.manager.Get(constants.JobType(jobType))

	if !exists {
		log.Println("No handler found for:", job.Type)
		return
	}

	handler.Handle(job)

	log.Println("Job completed:", job.ID)
}
