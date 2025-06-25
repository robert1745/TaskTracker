package main

import (
	"TaskTracker/task"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [add|update|delete|...]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <task_description>")
			return
		}
		desc := os.Args[2]
		if err := task.AddTask(desc); err != nil {
			fmt.Println("Error adding task:", err)
			return
		}

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		if err := task.ListTasks(filter); err != nil {
			fmt.Println("Error listing tasks", err)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <task_id> <new_description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task id ")
			return
		}
		newDesc := os.Args[3]
		if err := task.UpdateTask(id, newDesc); err != nil {
			fmt.Println("Error updating task:", err)
		} else {
			fmt.Println("Task updated successfully.")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:")
			return
		}
		if err = task.DeleteTask(id); err != nil {
			fmt.Println("Error deleting task : ", err)
		} else {
			fmt.Println("Task deleted successfully.")
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err := task.MarkTaskStatus(id, task.StatusDone); err != nil {
			fmt.Println("Error:", err)
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err := task.MarkTaskStatus(id, task.StatusInProgress); err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
