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

var InProgressList []*Task
var CompletedList []*Task
var ToDoList []*Task

func New(id int, description string, status string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func firstFreeID() int {
	//Переделать основной список в map[int:Task] и итерировать до первого ненахода числа в мапе?
	return 1
}

func Add(description string, status string) *Task {
	newTask := New(firstFreeID(), description, status)
	ToDoList = append(ToDoList, newTask)
	return newTask
}

func (t *Task) Update(description string) {
	t.Description = description
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
}
