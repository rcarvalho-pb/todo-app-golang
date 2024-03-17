package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	todo_repository "github.com/rcarvalho-pb/todo-app-golang/internal/adapter/output/repository/todo"
	"github.com/rcarvalho-pb/todo-app-golang/internal/controller"
	"github.com/rcarvalho-pb/todo-app-golang/internal/service"
)

const TODO_RESOURCE = "todos"

var todoController controller.TodoController

func initTodoRoutes(db *sqlx.DB) []Route {
	todoController = *controller.NewTodoController(service.NewTodoService(todo_repository.NewTodoRepository(db)))
	return todoRoutes
}

var todoRoutes = []Route{
	{
		Uri:            fmt.Sprintf("/%s", TODO_RESOURCE),
		Method:         http.MethodPost,
		Function:       todoController.CreateTodo,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s", TODO_RESOURCE),
		Method:         http.MethodGet,
		Function:       todoController.FindAllTodos,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:         http.MethodPut,
		Function:       todoController.UpdateTodo,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:         http.MethodGet,
		Function:       todoController.FindTodoById,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:         http.MethodDelete,
		Function:       todoController.DeleteTodoById,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/users/{id}", TODO_RESOURCE),
		Method:         http.MethodGet,
		Function:       todoController.FindAllTodosByUserId,
		Authentication: false,
	},
}
