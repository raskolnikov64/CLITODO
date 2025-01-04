package main

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t Task) PrintTask() {
	status := "Pending"
	if t.Completed {
		status = "Completed"
	}
	fmt.Printf("ID: %d | Description: %s | Status: %s\n", t.ID, t.Description, status)
}
