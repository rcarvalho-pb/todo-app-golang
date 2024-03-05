package controller

import (
	"io"
	"net/http"

	"github.com/rcarvalho-pb/todo-app-golang/internal/service"
)

type TodoController struct {
	*service.TodoService
}

func New(service *service.TodoService) *TodoController {
	return &TodoController{
		TodoService: service,
	}
}

func (td *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {

	}
}

func (td *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindAll(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindById(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindAllByUserId(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) DeleteById(w http.ResponseWriter, r *http.Request) {}
