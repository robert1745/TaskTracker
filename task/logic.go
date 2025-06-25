package task

import (
	"fmt"
	"strings"
	"time"
)

func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newID := getNextID(tasks)

	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	if err := SaveTasks(tasks); err != nil {
		return err
	}
	fmt.Println("Task added successfully.")
	return nil
}

func ListTasks(filter string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	count := 0
	filter = strings.ToLower(filter)
	for _, task := range tasks {
		if filter == "" || strings.EqualFold(task.Status, filter) {
			fmt.Printf("[%d] (%s) %s\n", task.ID, task.Status, task.Description)
			count++
		}
	}
	if count == 0 {
		fmt.Println("No tasks found")
	}
	return nil
}

func UpdateTask(id int, newDesc string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}
	if !updated {
		return fmt.Errorf("Task with ID %d not found", id)
	}

	return SaveTasks(tasks)
}

func DeleteTask(id int) error {

	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	deleted := false
	newTasks := make([]Task, 0)
	for i := range tasks {
		if tasks[i].ID == id {
			deleted = true
			continue
		}
		newTasks = append(newTasks, tasks[i])
	}

	if !deleted {
		return fmt.Errorf("Task with ID %d not found", id)
	}
	return SaveTasks(newTasks)

}

func MarkTaskStatus(id int, status string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("Task with ID %d not found", id)
	}
	return SaveTasks(tasks)
}
