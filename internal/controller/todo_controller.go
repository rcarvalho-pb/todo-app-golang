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

func NewTodoController(service *service.TodoService) *TodoController {
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

	id, err := td.TodoService.SaveTodo(&todo)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response_json.JSON(w, http.StatusCreated, id)
}

func (td *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}
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

	todo.ID = id

	id, err = td.TodoService.UpdateTodo(&todo)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response_json.JSON(w, http.StatusOK, id)
}

func (td *TodoController) FindAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := td.TodoService.FindAllTodos()
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todos)
}

func (td *TodoController) FindTodoById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	todo, err := td.TodoService.FindTodoById(id)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todo)
}

func (td *TodoController) FindAllTodosByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	todos, err := td.TodoService.FindAllTodosByUserId(id)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todos)
}

func (td *TodoController) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = td.TodoService.DeleteTodoById(id); err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, nil)
}
