package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	StatusTodo       = "TODO"
	StatusInProgress = "IN_PROGRESS"
	StatusDone       = "DONE"
)

func getNextID(tasks []Task) int {
	maxID := 0

	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}
