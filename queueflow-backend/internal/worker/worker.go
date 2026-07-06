package worker

import (
	"log"

	"queueflow/internal/queue"
)

type Worker struct {
	queue queue.Queue
}

func NewWorker(q queue.Queue) *Worker {
	return &Worker{queue: q}
}

func (w *Worker) Process() {
	job, err := w.queue.Pop()
	if err != nil {
		log.Println("queue error:", err)
		return
	}

	log.Println("Processing job:", job.ID, job.Type)

	switch job.Type {
	case "email":
		log.Println("Sending email to:", job.Payload)

	case "image":
		log.Println("Processing image:", job.Payload)
	}

	log.Println("Job completed:", job.ID)
}
