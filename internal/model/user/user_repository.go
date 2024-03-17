package user_model

type IUserRepository interface {
	SaveUser(*UserModel) (int64, error)
	FindAllUsers() (*[]UserModel, error)
	FindUserById(id int64) (*UserModel, error)
	DeleteUserById(id int64) error
}
