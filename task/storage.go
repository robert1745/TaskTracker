package task

import (
	"encoding/json"
	"errors"
	"os"
)

const file = "tasks.json"

func LoadTasks() ([]Task, error) {

	data, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if len(data) == 0 {
		return []Task{}, nil
	}
	// json object -> go struct
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func SaveTasks(tasks []Task) error {
	// go struct -> json
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
	return err

}
