package data

import (
	"errors"
	"fmt"
	"github.com/danxuZhang/gotodo/src/modules"
)

// TodoList A list of to-do tasks.
type TodoList struct {
	list []modules.Task
}

// NewTodoList creates a new empty to-do list.
func NewTodoList() *TodoList {
	return &TodoList{[]modules.Task{}}
}

// Get returns a to-do task at a given index.
func (receiver *TodoList) Get(index int) modules.Task {
	return receiver.list[index]
}

// Contains check whether a to-do task is in the to-do list or not.
func (receiver *TodoList) Contains(todo modules.Task) bool {
	for _, item := range receiver.list {
		if item == todo {
			return true
		}
	}
	return false
}

// IndexOf finds the index of a given to-do task, returns -1 if not found.
func (receiver *TodoList) IndexOf(todo modules.Task) int {
	for i, item := range receiver.list {
		if item == todo {
			return i
		}
	}
	return -1
}

// AddTask Adds a to-do task to the to-do list.
// Returns (true, nil) if the to-do list is modified.
// Returns (false, error) a problem occurred.
func (receiver *TodoList) AddTask(todo modules.Task) (bool, error) {
	receiver.list = append(receiver.list, todo)
	return true, nil
}

// RemoveItem Removes an item from the to-do list.
// Returns (true, nil) if the item is successfully removed.
// Returns (false, error) if the item is not contained in the list.
func (receiver *TodoList) RemoveItem(todo modules.Task) (bool, error) {
	index := receiver.IndexOf(todo)
	if index == -1 {
		return false, errors.New("todo object not found")
	}
	return receiver.RemoveIndex(index)
}

// RemoveIndex Removes a task at the given index.
// Returns (true, nil) if the removal is successful.
// Returns (false, error) if the input index is out of bound.
func (receiver *TodoList) RemoveIndex(index int) (bool, error) {
	if index < 0 || index > len(receiver.list) {
		return false, errors.New("index out of bound")
	}
	receiver.list = append(receiver.list[:index], receiver.list[index+1:]...)
	return true, nil
}

// ToString Converts the list to a string.
func (receiver *TodoList) ToString() string {
	FORMAT := "%d. %s\n"
	var res = ""
	for i, item := range receiver.list {
		res += fmt.Sprintf(FORMAT, i, item.ToString())
	}
	return res
}
