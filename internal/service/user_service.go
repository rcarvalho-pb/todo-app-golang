package service

import (
	"log"
	"time"

	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
)

type UserService struct {
	user_model.IUserRepository
}

func NewUserService(userRepository user_model.IUserRepository) *UserService {
	return &UserService{
		IUserRepository: userRepository,
	}
}

func (us *UserService) SaveUser(user *user_model.UserModel) (int64, error) {
	return us.IUserRepository.SaveUser(user)
}

func (us *UserService) UpdateUser(user *user_model.UserModel) (int64, error) {
	loadUser, err := us.IUserRepository.FindUserById(user.ID)
	log.Println(loadUser)
	if err != nil {
		return 0, err
	}

	if user.FirstName == "" {
		loadUser.FirstName = user.FirstName
	}

	if user.LastName == "" {
		loadUser.LastName = user.LastName
	}

	if user.Email != "" {
		loadUser.Email = user.Email
	}

	if user.Password != "" {
		loadUser.Password = user.Password
	}

	loadUser.LastModifiedDate = time.Now()

	return us.IUserRepository.SaveUser(loadUser)
}

func (us *UserService) FindAllUsers() (*[]user_model.UserModel, error) {
	return us.IUserRepository.FindAllUsers()
}

func (us *UserService) FindUserById(id int64) (*user_model.UserModel, error) {
	return us.IUserRepository.FindUserById(id)
}

func (us *UserService) DeleteUserById(id int64) error {
	return us.IUserRepository.DeleteUserById(id)
}
