package queue

import "queueflow/internal/models"

type Queue interface {
	Push(job models.Job) error
	Pop() (models.Job, error)
}
