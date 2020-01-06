package main

import (
	"fmt"
)

type Task {
	ID int
	Detail string
	done bool
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func main() {
	task := NewTask(1, "buy the milk")
	task.Finish()
	fmt.Printf("%+v", task)
}