package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("App is starting...")
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		fmt.Println("Commands: add, list, delete, edit,complete")
		return
	}

	command := args[1]
	fmt.Println("Command received:", command)
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: todo add <description>")
			return
		}
		description := args[2]
		newTask := Task{
			ID:          len(tasks) + 1,
			Description: description,
			Completed:   false,
		}
		tasks = append(tasks, newTask)
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task:", err)
		} else {
			fmt.Println("Task added successfully!")
		}

	case "list":
		fmt.Println("Executing 'list' command...")
		if len(tasks) == 0 {
			fmt.Println("No tasks available.")
		}
		for _, task := range tasks {
			task.PrintTask()
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		newTasks := []Task{}
		for _, task := range tasks {
			if task.ID != id {
				newTasks = append(newTasks, task)
			}
		}
		if len(newTasks) == len(tasks) {
			fmt.Println("Task not found.")
		} else {
			SaveTasks(newTasks)
			fmt.Println("Task deleted successfully!")
		}

	case "complete":
		if len(args) < 3 {
			fmt.Println("Usage: todo complete <id>")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		updated := false
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Completed = true
				updated = true
				break
			}
		}
		if !updated {
			fmt.Println("Task not found.")
		} else {
			SaveTasks(tasks)
			fmt.Println("Task marked as completed!")
		}
	case "edit":
		if len(args) < 4 {
			fmt.Println("Usage: todo edit <id> <field> <value>")
			fmt.Println("Fields: description, status")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}

		field := args[3]
		value := args[4]
		updated := false

		for i, task := range tasks {
			if task.ID == id {
				switch field {
				case "description":
					tasks[i].Description = value
					updated = true
				case "status":
					if value == "completed" {
						tasks[i].Completed = true
					} else if value == "pending" {
						tasks[i].Completed = false
					} else {
						fmt.Println("Invalid status. Use 'completed' or 'pending'.")
						return
					}
					updated = true
				default:
					fmt.Println("Invalid field. Use 'description' or 'status'.")
					return
				}
				break
			}
		}

		if !updated {
			fmt.Println("Task not found.")
		} else {
			if err := SaveTasks(tasks); err != nil {
				fmt.Println("Error saving tasks:", err)
			} else {
				fmt.Println("Task updated successfully!")
			}
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, list, delete, edit, complete")
	}
}
