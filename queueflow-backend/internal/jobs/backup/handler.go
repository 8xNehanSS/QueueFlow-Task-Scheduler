package backup

import (
	"log"
	"queueflow/internal/constants"
	"queueflow/internal/models"
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

func (b *BackupJob) Handle(job models.Job) {
	b.processing = true
	b.count++

	log.Println("Starting backup job:", job.ID)

	// backup logic here

	b.processing = false
}

func (b *BackupJob) GetStatus() string {
	if b.processing {
		return constants.HandlerProcessing
	}
	return constants.HandlerIdle
}
