package main

import (
	"log"
	"net/http"
	config "queueflow/configs"
	"queueflow/internal/api/routes"
	"queueflow/internal/queue"
	"queueflow/internal/service"
)

func main() {

	cfg := config.Load()
	q := queue.NewRedisQueue(cfg.RedisURL)

	jobService := service.NewJobService(q)

	router := routes.SetupRoutes(jobService)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
