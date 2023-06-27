package data

import (
	"database/sql"
	"time"

	"github.com/kublick/goexpense/internal/validator"
)

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users (name, email, last_name, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at`

	args := []interface{}{user.Name, user.Email, user.LastName, user.Password}

	return m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (m UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(user.Email != "", "email", "must be provided")
	v.Check(user.Password != "", "password", "must be provided")
	// v.Check(ValidateEmailFormat(user.Email), "email", "must be valid")
}

type MockUserModel struct{}

func (m MockUserModel) Insert(user *User) error {
	return nil
}

func (m MockUserModel) Get(id int64) (*User, error) {
	return nil, nil
}
