package worker

import (
	"queueflow/internal/constants"
	"queueflow/internal/jobs/op"
)

type Manager struct {
	handlers map[constants.JobType]op.IJob
}

func NewManager() *Manager {
	return &Manager{
		handlers: make(map[constants.JobType]op.IJob),
	}
}

func (m *Manager) Register(name constants.JobType, handler op.IJob) {
	m.handlers[name] = handler
}

func (m *Manager) Get(name constants.JobType) (op.IJob, bool) {
	handler, ok := m.handlers[name]
	return handler, ok
}
