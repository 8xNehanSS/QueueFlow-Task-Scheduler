package main

import (
	"log"
	"net/http"
	"queueflow/internal/api/routes"
	"queueflow/internal/queue"
	"queueflow/internal/service"
)

func main() {
	q := queue.NewRedisQueue()

	jobService := service.NewJobService(q)

	router := routes.SetupRoutes(jobService)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
