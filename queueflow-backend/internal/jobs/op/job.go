package op

import "queueflow/internal/models"

type IJob interface {
	Handle(models.Job) error
	GetStatus() (int, string)
}
