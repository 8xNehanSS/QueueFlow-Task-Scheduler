package queue

import (
	"sync"

	"queueflow/internal/models"
)

type InMemoryQueue struct {
	jobs []models.Job
	mu   sync.Mutex
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		jobs: make([]models.Job, 0),
	}
}

func (q *InMemoryQueue) Push(job models.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.jobs = append(q.jobs, job)
}

func (q *InMemoryQueue) Pop() (models.Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.jobs) == 0 {
		return models.Job{}, false
	}

	job := q.jobs[0]
	q.jobs = q.jobs[1:]

	return job, true
}
