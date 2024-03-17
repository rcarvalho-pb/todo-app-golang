package user_repository

import (
	"github.com/jmoiron/sqlx"
	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) SaveUser(user *user_model.UserModel) (int64, error) {
	res, err := ur.NamedExec("INSERT INTO users (first_name, last_name, email, password) VALUES (:first_name, :last_name, :email, :password)", user)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) FindAllUsers() (*[]user_model.UserModel, error) {
	var users []user_model.UserModel
	if err := ur.Select(&users, "SELECT * FROM users WHERE avaliable = TRUE ORDER BY first_name"); err != nil {
		return &[]user_model.UserModel{}, err
	}
	return &users, nil
}

func (ur *UserRepository) FindUserById(id int64) (*user_model.UserModel, error) {
	var user user_model.UserModel
	if err := ur.Get(&user, "SELECT * FROM users WHERE id=$1 AND avaliable = TRUE", id); err != nil {
		return &user_model.UserModel{}, err
	}

	return &user, nil
}

func (ur *UserRepository) FindAllUsersByTodoId(id int64) (*[]user_model.UserModel, error) {
	var users []user_model.UserModel
	if err := ur.Select(&users, "SELECT u.* FROM users u JOIN todos_users tu ON u.id = tu.user_id WHERE tu.todo_id=$1 AND u.avaliable = TRUE ORDER BY u.first_name", id); err != nil {
		return &[]user_model.UserModel{}, err
	}
	return &users, nil
}

func (ur *UserRepository) DeleteUserById(id int64) error {
	tx := ur.MustBegin()
	tx.MustExec("UPDATE users SET avaliable = FALSE WHERE id = $1", id)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
