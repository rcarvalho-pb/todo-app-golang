package controller

import (
	"encoding/json"
	"io"
	"net/http"

	todo_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/todo"
	response_json "github.com/rcarvalho-pb/todo-app-golang/internal/response"
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
		response_json.JSON(w, http.StatusBadRequest, err)
		return
	}

	var todo todo_model.TodoModel
	if err = json.Unmarshal(body, &todo); err != nil {
		response_json.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := td.TodoService.Save(&todo)
	if err != nil {
		response_json.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	response_json.JSON(w, http.StatusCreated, id)
}

func (td *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindAll(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindById(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) FindAllByUserId(w http.ResponseWriter, r *http.Request) {}

func (td *TodoController) DeleteById(w http.ResponseWriter, r *http.Request) {}
