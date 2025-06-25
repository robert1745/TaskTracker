# 🧠 Task Tracker CLI (Go)

A simple command-line task manager written in Go. Easily track, update, and manage your tasks from the terminal. Tasks are saved in a JSON file locally.

---

## 📦 Features

- Add new tasks with a description
- Update existing tasks
- Delete tasks
- Mark tasks as `TODO`, `IN_PROGRESS`, or `DONE`
- List all tasks or filter by status
- Persistent storage using a local `tasks.json` file

---

## 🏗 Folder Structure
``` bash
TaskTracker/
├── main.go # CLI entry point
├── go.mod # Go module file
├── task/ # Task logic and utilities
│ ├── logic.go # Business logic (add, update, delete, etc.)
│ ├── storage.go # JSON read/write helpers
│ └── task.go # Task model and constants
└── README.md # This file
```

###  How to run the project : 

```bash
go run main.go add "Buy groceries"
go run main.go list
go run main.go update 1 "Buy groceries and cook dinner"
go run main.go mark-in-progress 1
go run main.go mark-done 1
go run main.go delete 1
go run main.go list done
```

Project challenge : https://roadmap.sh/projects/task-tracker
