package worker

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/queue"
	"queueflow/internal/repository"
)

type Worker struct {
	queue   queue.Queue
	manager *Manager
	repo    *repository.JobRepository
}

func NewWorker(q queue.Queue, m *Manager, repo *repository.JobRepository) *Worker {
	return &Worker{
		queue:   q,
		manager: m,
		repo:    repo,
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

	w.repo.UpdateStatus(
		job.ID,
		"processing",
	)

	err = handler.Handle(job)
	if err != nil {
		log.Println("Error occurred while handling job:", err)
		w.repo.UpdateStatus(job.ID, "failed")
		return
	}

	w.repo.UpdateStatus(job.ID, "completed")
	log.Println("Job completed:", job.ID)
}
