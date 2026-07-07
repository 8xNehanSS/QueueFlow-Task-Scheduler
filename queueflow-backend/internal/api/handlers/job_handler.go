package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"queueflow/internal/models"
	"queueflow/internal/repository"
	"queueflow/internal/service"

	"github.com/google/uuid"
)

type JobHandler struct {
	service *service.JobService
	db      *sql.DB
}

func NewJobHandler(s *service.JobService, db *sql.DB) *JobHandler {
	return &JobHandler{service: s, db: db}
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	repo := repository.NewJobRepository(h.db)

	job := models.Job{
		ID:      uuid.New(),
		Type:    req.Type,
		Payload: req.Payload,
		Status:  "queued",
	}

	err := repo.Create(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	job, err = h.service.CreateJob(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(job)
}

func (h *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewJobRepository(h.db)
	jobs, err := repo.List(nil, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jobs)
}
