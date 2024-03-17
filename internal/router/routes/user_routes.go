package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	user_repository "github.com/rcarvalho-pb/todo-app-golang/internal/adapter/output/repository/user"
	"github.com/rcarvalho-pb/todo-app-golang/internal/controller"
	"github.com/rcarvalho-pb/todo-app-golang/internal/service"
)

const USER_RESOURCE = "users"

var userController controller.UserController

func initUserRoutes(db *sqlx.DB) []Route {
	userController = *controller.NewUserController(service.NewUserService(user_repository.NewUserRepository(db)))
	return userRoutes
}

var userRoutes = []Route{
	{
		Uri:            fmt.Sprintf("/%s", USER_RESOURCE),
		Method:         http.MethodGet,
		Function:       userController.FindAllUsers,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s", USER_RESOURCE),
		Method:         http.MethodPost,
		Function:       userController.CreateUser,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:         http.MethodGet,
		Function:       userController.FindUserById,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:         http.MethodPut,
		Function:       userController.UpdateUser,
		Authentication: false,
	},
	{
		Uri:            fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:         http.MethodDelete,
		Function:       userController.DeleteUserById,
		Authentication: false,
	},
}
