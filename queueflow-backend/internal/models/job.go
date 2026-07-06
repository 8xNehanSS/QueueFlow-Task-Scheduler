package models

import "time"

type Job struct {
	ID        string
	Type      string
	Payload   string
	Status    string
	CreatedAt time.Time
}
