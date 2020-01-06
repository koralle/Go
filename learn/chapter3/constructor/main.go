package main

import (
	"fmt"
)

type Task struct {
	ID int
	Detail string
	done bool
}

func NewTask(id int, detail string) *Task {
	task := &Task {
		ID: id,
		Detail: detail,
		done: false,
	}
	return task
}

func main(){
	task := NewTask(1, "buy the milk")
	fmt.Printf("%+v", task)
}