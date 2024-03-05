package todo_model

import (
	"errors"
	"time"

	user_model "github.com/rcarvalho-pb/todo-app-golang/internal/model/user"
)

type TodoModel struct {
	ID               int64             `db:"id,omitempty"`
	Name             string            `db:"name,omitempty"`
	Description      string            `db:"description,omitempty"`
	Status           int64             `db:"status,omitempty"`
	CreatedAt        time.Time         `db:"created_at,omitempty"`
	LastModifiedDate time.Time         `db:"last_modified_date,omitempty"`
	Users            []user_model.User `db:"-"`
}

var StatusType = map[string]int64{
	"To Do": 1,
	"Doing": 2,
	"Done":  3,
}

func GetStatusString(value int64) (string, error) {
	for k := range StatusType {
		if StatusType[k] == value {
			return k, nil
		}
	}
	return "", errors.New("not a valid status int")
}
