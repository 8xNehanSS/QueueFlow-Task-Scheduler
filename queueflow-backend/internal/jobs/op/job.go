package op

import "queueflow/internal/models"

type IJob interface {
	Handle(models.Job)
	GetStatus() string
}
