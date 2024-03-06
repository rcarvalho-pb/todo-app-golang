package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var todo todo_model.TodoModel
	if err = json.Unmarshal(body, &todo); err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := td.TodoService.Save(&todo)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response_json.JSON(w, http.StatusCreated, id)
}

func (td *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var todo todo_model.TodoModel
	if err = json.Unmarshal(body, &todo); err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	id, err := td.TodoService.UpdateTodo(&todo)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response_json.JSON(w, http.StatusOK, id)
}

func (td *TodoController) FindAll(w http.ResponseWriter, r *http.Request) {
	todos, err := td.TodoService.FindAll()
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todos)
}

func (td *TodoController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	todo, err := td.TodoService.FindById(id)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todo)
}

func (td *TodoController) FindAllByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	todos, err := td.TodoService.FindAllByUserId(id)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todos)
}

func (td *TodoController) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = td.TodoService.DeleteById(id); err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, nil)
}
