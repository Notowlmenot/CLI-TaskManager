package main

import (
	tasks "CLI-TaskManager/Tasks"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: Была задумка обработки команд через Map, получая функцию чрез ключ вводимой команды в строке
// var commands = map[string]func(){
// "add": add(),
// "update":  update,
// "deletemark-in-progress":  showHelp,
// "mark-done": _,
// "list": _,
// }

// # Adding a new task
// task-cli add "Buy groceries"
// # Output: Task added successfully (ID: 1)

// # Updating and deleting tasks
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # Marking a task as in progress or done
// task-cli mark-in-progress 1
// task-cli mark-done 1

// # Listing all tasks
// task-cli list

// # Listing tasks by status
// task-cli list done
// task-cli list todo
// task-cli list in-progress

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		UserInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при команды: ", err)
		}
		input := strings.Fields(UserInput)

		command := ""
		if len(input) > 0 {
			command = input[0]
		}

		switch command {
		case "add":
			if len(input) >= 2 {
				description := strings.Join(input[1:], " ")
				newTask := tasks.Add(description, tasks.TODO)
				fmt.Println("Успешно добавлена задача №", newTask.ID)
			} else {
				fmt.Println(input[1:])
				fmt.Println("Необходимо указать описание задачи.")
				continue
			}

		case "delete":
			TaskNum, err := (strconv.Atoi(input[1])) // -1
			if err != nil {
				println("Неверно введен номер задачи.")
				continue
			}
			tasks.Delete(TaskNum)
		case "update":
			if len(input) >= 3 {
				TaskNum, err := (strconv.Atoi(input[1])) // -1
				newDescription := strings.Join(input[1:], " ")
				if err != nil {
					println("Неверно введен номер задачи.")
				}
				tasks.Update(TaskNum, newDescription)
			} else {
				fmt.Println("Неверный ввод команды. Необходимые аргументы: update NUM NEW_DESCRIPTION")
			}
		case "list":
			if len(input) >= 2 {
				switch input[1] {
				case "done":
					for key, value := range tasks.InProgressList {
						fmt.Println("Список выполненых задач: \n")
						fmt.Printf("Задача №%d: %s \n", key, value.Description)
					}
				case "all":
					for key, value := range tasks.ToDoList {
						fmt.Printf("Список всех задач: \n")
						fmt.Printf("Задача №%d: %s. Статус: %s \n", key, value.Description, value.Status)
					}
				case "in-progress":
					for key, value := range tasks.InProgressList {
						fmt.Println("Список задач в процессе выполнения: \n")
						fmt.Printf("Задача №%d: %s \n", key, value.Description)
					}
				default:
					println("Неверно введен тип списка. Возможные варианты: \n list todo \n list done \n list in-progress \n")
				}
			} else {
				fmt.Println("Неверно введен тип списка. Возможные варианты: \n list todo \n list done \n list in-progress \n")
			}
		case "mark-in-progress":
			if len(input) >= 2 {
				TaskNum, err := (strconv.Atoi(input[1])) // -1
				if err != nil {
					println("Неверно введен номер задачи.")
				}
				tasks.UpdateStatus(TaskNum, tasks.INPROGRESS)
			}
		case "mark-done":
			if len(input) >= 2 {
				TaskNum, err := (strconv.Atoi(input[1])) // -1
				if err != nil {
					println("Неверно введен номер задачи.")
				}
				tasks.UpdateStatus(TaskNum, tasks.COMPLETED)
			}
		case "help":
			fmt.Println("Список команд:")
		default:
			//
		}
	}
}
