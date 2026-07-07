package main

import (
	"log"
	"net/http"
	config "queueflow/configs"
	"queueflow/internal/api/routes"
	"queueflow/internal/database"
	"queueflow/internal/queue"
	"queueflow/internal/service"
)

func main() {

	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	q := queue.NewRedisQueue(cfg.RedisURL)

	jobService := service.NewJobService(q)

	router := routes.SetupRoutes(jobService, db)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
