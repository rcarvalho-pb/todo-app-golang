package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/todo-app-golang/internal/adapter/output/repository"
	"github.com/rcarvalho-pb/todo-app-golang/internal/controller"
	"github.com/rcarvalho-pb/todo-app-golang/internal/service"
)

const TODO_RESOURCE = "todos"

var todoController controller.TodoController

func initTodoRoutes(db *sqlx.DB) []Route {
	todoController = *controller.New(service.New(repository.New(db)))
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
		Function:       todoController.FindAll,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:         http.MethodGet,
		Function:       todoController.FindById,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:         http.MethodDelete,
		Function:       todoController.DeleteById,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/users/{id}", TODO_RESOURCE),
		Method:         http.MethodGet,
		Function:       todoController.FindAllByUserId,
		Authentication: false,
	},
}
