package todo_repository

import (
	"github.com/jmoiron/sqlx"
	todo_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/todo"
	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
)

type TodoRepository struct {
	*sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (tr *TodoRepository) SaveTodo(todo *todo_model.TodoModel) (int64, error) {
	res, err := tr.NamedExec("INSERT INTO todos (name, description) VALUES (:name, :description)", todo)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if todo.Users != nil {
		for _, u := range todo.Users {
			_, err = tr.Exec("INSERT INTO todos_users VALUES ($1, $2)", id, u.ID)
			if err != nil {
				return id, err
			}
		}
	}

	return id, nil
}

func (tr *TodoRepository) FindAllTodos() (*[]todo_model.TodoModel, error) {
	var todos []todo_model.TodoModel
	if err := tr.Select(&todos, "SELECT * FROM todos WHERE avaliable = TRUE ORDER BY name"); err != nil {
		return &[]todo_model.TodoModel{}, err
	}
	for i, todo := range todos {
		var users []user_model.UserModel
		if err := tr.Select(&users, "SELECT u.* FROM users u JOIN todos_users tu ON u.id = tu.user_id WHERE tu.todo_id=$1 AND u.avaliable = TRUE ORDER BY u.first_name", todo.ID); err != nil {
			return &[]todo_model.TodoModel{}, err
		}
		todos[i].Users = users
	}

	return &todos, nil
}

func (tr *TodoRepository) FindTodoById(id int64) (*todo_model.TodoModel, error) {
	var todo todo_model.TodoModel
	if err := tr.Get(&todo, "SELECT * FROM todos WHERE id=$1 AND avaliable = TRUE", id); err != nil {
		return &todo_model.TodoModel{}, err
	}
	var users []user_model.UserModel
	if err := tr.Select(&users, "SELECT u.* FROM users u JOIN todos_users tu ON u.id = tu.user_id WHERE tu.todo_id=$1", todo.ID); err != nil {
		return &todo_model.TodoModel{}, err
	}
	todo.Users = users

	return &todo, nil
}

func (tr *TodoRepository) FindAllTodosByUserId(id int64) (*[]todo_model.TodoModel, error) {
	var todos []todo_model.TodoModel
	if err := tr.Select(&todos, "SELECT t.* FROM todos t JOIN todos_users tu ON t.id = tu.todo_id WHERE tu.user_id=$1 AND t.avaliable = TRUE ORDER BY t.name", id); err != nil {
		return &[]todo_model.TodoModel{}, err
	}

	for i, todo := range todos {
		var users []user_model.UserModel
		if err := tr.Select(&users, "SELECT u.* FROM users u JOIN todos_users tu ON u.id = tu.user_id WHERE u.avaliable = TRUE AND tu.todo_id=$1", todo.ID); err != nil {
			return &[]todo_model.TodoModel{}, err
		}
		todos[i].Users = users
	}
	return &todos, nil
}

func (tr *TodoRepository) DeleteTodoById(id int64) error {
	tx := tr.MustBegin()
	tx.MustExec("UPDATE todos SET avaliable = FALSE WHERE id = $1", id)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
