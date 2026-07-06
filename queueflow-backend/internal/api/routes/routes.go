package routes

import (
	"net/http"

	"queueflow/internal/api/handlers"
	"queueflow/internal/service"
)

func SetupRoutes(jobService *service.JobService) http.Handler {
	mux := http.NewServeMux()

	jobHandler := handlers.NewJobHandler(jobService)

	mux.HandleFunc("/jobs", jobHandler.CreateJob)

	return mux
}
