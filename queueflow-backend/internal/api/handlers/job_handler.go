package handlers

import (
	"encoding/json"
	"net/http"

	"queueflow/internal/service"
)

type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(s *service.JobService) *JobHandler {
	return &JobHandler{service: s}
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	job, err := h.service.CreateJob(req.Type, req.Payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(job)
}
