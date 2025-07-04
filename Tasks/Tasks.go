package tasks

import (
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(id int, description string, status string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) Delete() {
	t = nil
}

func (t *Task) Update(description string) {
	t.Description = description
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
}
