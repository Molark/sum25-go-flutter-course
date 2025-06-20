package taskmanager

import (
	"errors"
	"time"
)

var (
	// ErrTaskNotFound is returned when a task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrEmptyTitle is returned when the task title is empty
	ErrEmptyTitle = errors.New("task title cannot be empty")
	// ErrInvalidID is returned when the task ID is invalid
	ErrInvalidID = errors.New("invalid task ID")
)

// Task represents a single task
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// TaskManager manages a collection of tasks
type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	// TODO: Implement task manager initialization
	newTaskManager := TaskManager{
		tasks:  map[int]*Task{},
		nextID: 1,
	}
	return &newTaskManager
}

// AddTask adds a new task to the manager
func (tm *TaskManager) AddTask(title, description string) (*Task, error) {
	// TODO: Implement task addition
	if len(title) == 0 {
		return nil, ErrEmptyTitle
	}
	newTask := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	tm.tasks[tm.nextID] = &newTask
	tm.nextID = tm.nextID + 1
	return &newTask, nil

}

// UpdateTask updates an existing task
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	// TODO: Implement task update
	if id < 1 {
		return ErrInvalidID
	}
	task, ok := tm.tasks[id]
	if ok == false {
		return ErrTaskNotFound
	}
	if len(title) == 0 {
		return ErrEmptyTitle
	}
	task.Title = title
	task.Description = description
	task.Done = done
	return nil
}

// DeleteTask removes a task from the manager
func (tm *TaskManager) DeleteTask(id int) error {
	// TODO: Implement task deletion
	if id < 1 {
		return ErrInvalidID
	}
	_, ok := tm.tasks[id]
	if ok == false {
		return ErrTaskNotFound
	}
	delete(tm.tasks, id)
	return nil
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	// TODO: Implement task retrieval
	if id < 1 {
		return nil, ErrInvalidID
	}
	task, ok := tm.tasks[id]
	if ok == false {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

// ListTasks returns all tasks, optionally filtered by done status
func (tm *TaskManager) ListTasks(filterDone *bool) []*Task {
	// TODO: Implement task listing with optional filter
	tasks := make([]*Task, 0)
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
