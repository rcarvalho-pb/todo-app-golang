package service

import (
	"time"

	todo_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/todo"
)

type TodoService struct {
	todo_model.ITodoRepository
}

func NewTodoService(todoRepository todo_model.ITodoRepository) *TodoService {
	return &TodoService{
		ITodoRepository: todoRepository,
	}
}

func (ts *TodoService) CreateTodo(todo *todo_model.TodoModel) (int64, error) {
	return ts.SaveTodo(todo)
}

func (ts *TodoService) UpdateTodo(todo *todo_model.TodoModel) (int64, error) {
	loadTodo, err := ts.FindTodoById(todo.ID)
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

	return ts.ITodoRepository.SaveTodo(loadTodo)
}

func (ts *TodoService) FindAllTodos() (*[]todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindAllTodos()
}

func (ts *TodoService) FindTodoById(id int64) (*todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindTodoById(id)
}

func (ts *TodoService) FindAllTodosByUserId(id int64) (*[]todo_model.TodoModel, error) {
	return ts.ITodoRepository.FindAllTodosByUserId(id)
}

func (ts *TodoService) DeleteTodoById(id int64) error {
	return ts.ITodoRepository.DeleteTodoById(id)
}
