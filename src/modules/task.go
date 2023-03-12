package modules

import (
	"errors"
	"fmt"
	"time"
)

// Task Represents a To-Do Task.
// Contains whether the task is done or not, text description and due time.
type Task struct {
	isDone      bool
	description string
	dueTime     time.Time
}

// NewTask Creates a new not-done Task.
func NewTask(description string, dueTime time.Time) *Task {
	task := Task{false, description, dueTime}
	return &task
}

// IsDone Checks whether a task is done or not.
func (receiver *Task) IsDone() bool {
	return receiver.isDone
}

// Mark Marks a task to be done.
// Returns (true, nil) if the action is successful.
// Returns (false, error) if the task is already marked.
func (receiver *Task) Mark() (bool, error) {
	if receiver.isDone {
		return false, errors.New("task is already marked")
	}
	receiver.isDone = true
	return true, nil
}

// Unmark Marks a task to be not done yet.
// Returns (true, nil) if the action is successful.
// Returns (false, error) if the task is already not marked.
func (receiver *Task) Unmark() (bool, error) {
	if !receiver.isDone {
		return false, errors.New("task is already not marked")
	}
	receiver.isDone = false
	return true, nil
}

// ToString converts the task to a string.
func (receiver *Task) ToString() string {
	FORMAT := "[%s] %s (due: %s)"
	var status string
	if receiver.isDone {
		status = "x"
	} else {
		status = " "
	}
	return fmt.Sprintf(FORMAT, status, receiver.description, receiver.dueTime)
}
