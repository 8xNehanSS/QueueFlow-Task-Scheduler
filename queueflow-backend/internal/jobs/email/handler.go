package email

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/models"
	"queueflow/internal/utils"
	"time"
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

func (e *EmailJob) Handle(job models.Job) error {
	e.processing = true
	e.count++

	log.Println("Starting email job:", job.ID)

	// email logic here
	sleepTimeDuration := time.Duration(utils.RandomNumber(1, 5)) * time.Second
	time.Sleep(sleepTimeDuration)

	e.processing = false
	return nil
}

func (e *EmailJob) GetStatus() (int, string) {
	if e.processing {
		return e.count, constants.HandlerProcessing
	}
	return e.count, constants.HandlerIdle
}
