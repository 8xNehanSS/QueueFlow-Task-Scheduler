package email

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/models"
)

type EmailJob struct {
	count      int
	processing bool
}

func NewJobHandler() *EmailJob {
	return &EmailJob{
		count:      0,
		processing: false,
	}
}

func (e *EmailJob) Handle(job models.Job) {
	e.processing = true
	e.count++

	log.Println("Starting email job:", job.ID)

	// email logic here

	e.processing = false
}

func (e *EmailJob) GetStatus() string {
	if e.processing {
		return constants.HandlerProcessing
	}
	return constants.HandlerIdle
}
