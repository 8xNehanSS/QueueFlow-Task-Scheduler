package service

import (
	"queueflow/internal/models"
	"queueflow/internal/queue"
)

type JobService struct {
	queue queue.Queue
}

func NewJobService(q queue.Queue) *JobService {
	return &JobService{queue: q}
}

func (s *JobService) CreateJob(job models.Job) (models.Job, error) {

	err := s.queue.Push(job)
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}
