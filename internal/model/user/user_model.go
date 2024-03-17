package user_model

import (
	"errors"
	"time"
)

type UserModel struct {
	ID               int64     `db:"id,omitempty"`
	FirstName        string    `db:"first_name,omitempty" json:"first_name"`
	LastName         string    `db:"last_name,omitempty" json:"last_name"`
	Email            string    `db:"email,omitempty" json:"email"`
	Password         string    `json:"password"`
	Role             string    `db:"role" json:"role"`
	CreatedAt        time.Time `db:"created_at,omitempty" json:"created_at"`
	LastModifiedDate time.Time `db:"last_modified_date,omitempty" json:"last_modified_date"`
	Avaliable        bool      `db:"avaliable" json:"avaliable"`
}

var UserRole = map[string]int64{
	"user":  1,
	"admin": 2,
}

func GetStatusString(value int64) (string, error) {
	for k := range UserRole {
		if UserRole[k] == value {
			return k, nil
		}
	}
	return "", errors.New("not a valid user role int")
}
