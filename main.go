package main

import (
	tasks "CLI-TaskManager/Tasks"
	"fmt"
)

func main() {
	Task1 := tasks.New(1, "", "")
	fmt.Println(Task1)
}
