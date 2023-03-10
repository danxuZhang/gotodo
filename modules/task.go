package modules

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	isDone      bool
	description string
	dueTime     time.Time
}

func NewTask(description string, dueTime time.Time) *Task {
	task := Task{false, description, dueTime}
	return &task
}

func (receiver *Task) IsDone() bool {
	return receiver.isDone
}

func (receiver *Task) Mark() (bool, error) {
	if receiver.isDone {
		return false, errors.New("task is already marked")
	}
	receiver.isDone = true
	return true, nil
}

func (receiver *Task) Unmark() (bool, error) {
	if !receiver.isDone {
		return false, errors.New("task is already not marked")
	}
	receiver.isDone = false
	return true, nil
}

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
