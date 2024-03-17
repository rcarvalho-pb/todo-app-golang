package todo_model

type ITodoRepository interface {
	SaveTodo(*TodoModel) (int64, error)
	FindAllTodos() (*[]TodoModel, error)
	FindTodoById(id int64) (*TodoModel, error)
	FindAllTodosByUserId(id int64) (*[]TodoModel, error)
	DeleteTodoById(id int64) error
}
