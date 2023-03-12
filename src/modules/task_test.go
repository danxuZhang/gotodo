package modules

import (
	"testing"
	"time"
)

func TestConstructor(t *testing.T) {
	// task := NewTodo("task", time.Now())
	task := Task{false, "task", time.Now()}
	if task.IsDone() {
		t.Error("default isDone should be false")
	}
}

func TestTask_Mark(t *testing.T) {
	task := NewTask("eat lunch", time.Now())
	if task.IsDone() {
		t.Error("default isDone error")
	}
	_, err := task.Mark()
	if err != nil || !task.IsDone() {
		t.Error("mark task error")
	}
	_, err = task.Mark()
	if err == nil {
		t.Error("duplicate mark task error")
	}
}

func TestTask_Unmark(t *testing.T) {
	task := NewTask("eat lunch", time.Now())
	if task.IsDone() {
		t.Error("default isDone error")
	}
	_, err := task.Unmark()
	if err == nil {
		t.Error("mark task error")
	}
	_, err = task.Mark()
	if !task.isDone {
		t.Error("mark error")
	}
	_, err = task.Unmark()
	if err != nil || task.IsDone() {
		t.Error("duplicate mark task error")
	}
}
