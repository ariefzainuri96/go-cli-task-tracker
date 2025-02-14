package main

import (
	"encoding/json"
	"fmt"
)

type TaskStatus string

const (
	TODO        = "todo"
	IN_PROGRESS = "in-progress"
	DONE        = "done"
)

type Task struct {
	Id          int
	Description string
	Status      TaskStatus
	CreatedAt   string
	UpdatedAt   string
}

func (t Task) ToJson() {
	value, err := json.Marshal(t)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(value))
}

type Command struct {
	command string
	example string
}
