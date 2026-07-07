package backup

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/models"
	"queueflow/internal/utils"
	"time"
)

type BackupJob struct {
	count      int
	processing bool
}

func NewJobHandler() *BackupJob {
	return &BackupJob{
		count:      0,
		processing: false,
	}
}

func (b *BackupJob) Handle(job models.Job) error {
	b.processing = true
	b.count++

	log.Println("Starting backup job:", job.ID)

	// backup logic here
	sleepTimeDuration := time.Duration(utils.RandomNumber(1, 5)) * time.Second
	time.Sleep(sleepTimeDuration)

	b.processing = false
	return nil
}

func (b *BackupJob) GetStatus() (int, string) {
	if b.processing {
		return b.count, constants.HandlerProcessing
	}
	return b.count, constants.HandlerIdle
}
