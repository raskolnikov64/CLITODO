package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const dataFile = "./tasks.json"

func LoadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File not found, returning empty tasks.")
			return []Task{}, nil
		}
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println("Tasks loaded successfully:", tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	} else {
		fmt.Println("Tasks saved successfully.")
	}
	return err
}
