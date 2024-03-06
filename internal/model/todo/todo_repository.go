package todo_model

type ITodoRepository interface {
	Save(*TodoModel) (int64, error)
	FindAll() (*[]TodoModel, error)
	FindById(id int64) (*TodoModel, error)
	FindAllByUserId(id int64) (*[]TodoModel, error)
	DeleteById(id int64) error
}
