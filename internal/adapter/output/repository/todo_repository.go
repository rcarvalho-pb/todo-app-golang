package repository

import (
	"github.com/jmoiron/sqlx"
	todo_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/todo"
	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
)

type TodoRepository struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (tr *TodoRepository) Save(todo *todo_model.TodoModel) (int64, error) {
	res, err := tr.NamedExec("INSERT INTO todos (name, description) VALUES (:name, :description)", todo)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TodoRepository) FindAll() (*[]todo_model.TodoModel, error) {
	var todos []todo_model.TodoModel
	if err := tr.Select(&todos, "SELECT * FROM todos ORDER BY name"); err != nil {
		return &[]todo_model.TodoModel{}, err
	}

	return &todos, nil
}

func (tr *TodoRepository) FindById(id int64) (*todo_model.TodoModel, error) {
	var todo todo_model.TodoModel
	if err := tr.Get(&todo, "SELECT * FROM todos WHERE id=$1", id); err != nil {
		return &todo_model.TodoModel{}, err
	}

	return &todo, nil
}

func (tr *TodoRepository) FindAllByUserId(id int64) (*[]todo_model.TodoModel, error) {
	var todos []todo_model.TodoModel
	if err := tr.Select(&todos, "SELECT t.* FROM todos t JOIN todos_users tu ON t.id = tu.todo_id WHERE tu.user_id=$1 ORDER BY t.name", id); err != nil {
		return &[]todo_model.TodoModel{}, err
	}

	for i, todo := range todos {
		var users []user_model.User
		if err := tr.Select(&users, "SELECT u.* FROM users u JOIN todos_users tu ON u.id = tu.user_id WHERE tu.todo_id=$1", todo.ID); err != nil {
			return &[]todo_model.TodoModel{}, err
		}
		todos[i].Users = users
	}
	return &todos, nil
}

func (tr *TodoRepository) DeleteById(id int64) (bool, error) {
	tx := tr.MustBegin()
	tx.MustExec("DELETE FROM todos_users WHERE todo_id=$1", id)
	tx.MustExec("DELETE FROM todos WHERE id=$1", id)
	if err := tx.Commit(); err != nil {
		return false, err
	}

	return true, nil
}
