package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
	response_json "github.com/rcarvalho-pb/todo-app-golang/internal/response"
	"github.com/rcarvalho-pb/todo-app-golang/internal/service"
)

type UserController struct {
	*service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user user_model.UserModel
	if err = json.Unmarshal(body, &user); err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := uc.UserService.SaveUser(&user)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response_json.JSON(w, http.StatusCreated, id)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var user user_model.UserModel
	if err = json.Unmarshal(body, &user); err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	id, err := uc.UserService.UpdateUser(&user)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response_json.JSON(w, http.StatusOK, id)
}

func (uc *UserController) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.FindAllUsers()
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, users)
}

func (uc *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	todo, err := uc.UserService.FindUserById(id)
	if err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, todo)
}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response_json.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = uc.UserService.DeleteUserById(id); err != nil {
		response_json.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response_json.JSON(w, http.StatusOK, nil)
}
