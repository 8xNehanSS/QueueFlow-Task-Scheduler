package main

import (
	"log"
	config "queueflow/configs"
	"queueflow/internal/constants"
	"queueflow/internal/jobs/backup"
	"queueflow/internal/jobs/email"
	"queueflow/internal/queue"
	"queueflow/internal/worker"
)

func main() {

	cfg := config.Load()
	q := queue.NewRedisQueue(cfg.RedisURL)

	manager := worker.NewManager()
	manager.Register(constants.JobBackup, backup.NewJobHandler())
	manager.Register(constants.JobEmail, email.NewJobHandler())

	log.Println("QueueFlow Worker")
	workerPool := worker.NewWorkerPool(q, cfg.WorkerCount, manager)
	workerPool.Start()
}
