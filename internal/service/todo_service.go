package service

import (
	"time"

	todo_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/todo"
)

type TodoService struct {
	todo_model.ITodoRepository
}

func New(todoRepository todo_model.ITodoRepository) *TodoService {
	return &TodoService{
		ITodoRepository: todoRepository,
	}
}

func (ts *TodoService) CreateTodo(todo *todo_model.TodoModel) (int64, error) {
	return ts.Save(todo)
}

func (ts *TodoService) UpdateTodo(todo *todo_model.TodoModel) (int64, error) {
	loadTodo, err := ts.FindById(todo.ID)
	if err != nil {
		return 0, err
	}

	if todo.Name == "" {
		loadTodo.Name = todo.Name
	}

	if todo.Description != "" {
		loadTodo.Description = todo.Description
	}

	if todo.Status != 0 {
		loadTodo.Status = todo.Status
	}

	loadTodo.LastModifiedDate = time.Now()

	if todo.Users != nil {
		loadTodo.Users = todo.Users
	}

	return ts.ITodoRepository.Save(loadTodo)
}

func (ts *TodoService) FindAll() (*[]todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindAll()
}

func (ts *TodoService) FindById(id int64) (*todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindById(id)
}

func (ts *TodoService) FindAllByUserId(id int64) (*[]todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindAllByUserId(id)
}

func (ts *TodoService) DeleteById(id int64) error {
	return ts.ITodoRepository.DeleteById(id)
}
