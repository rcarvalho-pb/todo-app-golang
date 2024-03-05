package user_model

import "time"

type User struct {
	ID               int64  `db:"id,omitempty"`
	FirstName        string `db:"first_name,omitempty"`
	LastName         string `db:"last_name,omitempty"`
	Email            string `db:"email,omitempty"`
	Password         string
	CreatedAt        time.Time `db:"created_at,omitempty"`
	LastModifiedDate time.Time `db:"last_modified_date,omitempty"`
}
