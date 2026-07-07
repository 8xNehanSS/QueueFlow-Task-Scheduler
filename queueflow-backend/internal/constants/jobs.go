package constants

type JobType string

const (
	JobEmail        JobType = "email"
	JobBackup       JobType = "backup"
	JobReport       JobType = "report"
	JobImage        JobType = "image"
	JobNotification JobType = "notification"
)
