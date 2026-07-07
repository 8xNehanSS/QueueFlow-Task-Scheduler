package routes

import (
	"database/sql"
	"net/http"

	"queueflow/internal/api/handlers"
	"queueflow/internal/auth"
	"queueflow/internal/permissions"
	"queueflow/internal/service"
)

func SetupRoutes(jobService *service.JobService, db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	jobHandler := handlers.NewJobHandler(jobService, db)
	authHandler := handlers.NewAuthHandler(auth.NewAuthService(db))

	// Public routes
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	mux.Handle(
		"POST /jobs",
		auth.Protected(
			http.HandlerFunc(jobHandler.CreateJob),
			permissions.JOB_CREATE,
		),
	)

	mux.Handle(
		"GET /jobs",
		auth.Protected(
			http.HandlerFunc(jobHandler.GetJobs),
			permissions.JOB_VIEW,
		),
	)

	return mux
}
