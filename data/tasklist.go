package data

import (
	"errors"
	"fmt"
	"github.com/danxuZhang/gotodo/modules"
)

type TodoList struct {
	list []modules.Task
}

func NewTodoList() *TodoList {
	return &TodoList{[]modules.Task{}}
}

func (receiver *TodoList) Get(index int) modules.Task {
	return receiver.list[index]
}

func (receiver *TodoList) Contains(todo modules.Task) bool {
	for _, item := range receiver.list {
		if item == todo {
			return true
		}
	}
	return false
}

func (receiver *TodoList) IndexOf(todo modules.Task) int {
	for i, item := range receiver.list {
		if item == todo {
			return i
		}
	}
	return -1
}

func (receiver *TodoList) AddTask(todo modules.Task) (bool, error) {
	receiver.list = append(receiver.list, todo)
	return true, nil
}

func (receiver *TodoList) RemoveItem(todo modules.Task) (bool, error) {
	index := receiver.IndexOf(todo)
	if index == -1 {
		return false, errors.New("todo object not found")
	}
	return receiver.RemoveIndex(index)
}

func (receiver *TodoList) RemoveIndex(index int) (bool, error) {
	if index < 0 || index > len(receiver.list) {
		return false, errors.New("index out of bound")
	}
	receiver.list = append(receiver.list[:index], receiver.list[index+1:]...)
	return true, nil
}

func (receiver *TodoList) ToString() string {
	FORMAT := "%d. %s\n"
	var res = ""
	for i, item := range receiver.list {
		res += fmt.Sprintf(FORMAT, i, item.ToString())
	}
	return res
}
