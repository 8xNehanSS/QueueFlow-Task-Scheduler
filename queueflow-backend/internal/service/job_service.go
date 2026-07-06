package service

import (
	"github.com/google/uuid"

	"queueflow/internal/models"
	"queueflow/internal/queue"
)

type JobService struct {
	queue queue.Queue
}

func NewJobService(q queue.Queue) *JobService {
	return &JobService{queue: q}
}

func (s *JobService) CreateJob(jobType string, payload string) (models.Job, error) {

	job := models.Job{
		ID:      uuid.NewString(),
		Type:    jobType,
		Payload: payload,
		Status:  "queued",
	}

	err := s.queue.Push(job)
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}
