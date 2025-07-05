package tasks

import (
	"fmt"
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
var ToDoList = make(map[int]*Task)

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
	ToDoList[firstFreeID()] = newTask
	return newTask
}

func Delete(num int) {
	_, ok := ToDoList[num]
	if ok {
		delete(ToDoList, num)
		fmt.Printf("Задач №%d успешно удалена \n", num)
	} else {
		fmt.Println("Неверно введен номер задачи")
	}
}

func Update(num int, description string) {
	_, ok := ToDoList[num]
	if ok {
		ToDoList[num].Description = description
		fmt.Printf("Задач №%d успешно обновлена \n", num)
	} else {
		fmt.Println("Неверно введен номер задачи")
	}
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
}
