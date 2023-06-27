package data

import (
	"database/sql"
	"errors"
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

	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT id, created_at, updated_at, name, email, last_name, password
	FROM USERS
	WHERE id = $1`

	var user User

	err := m.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
		&user.Email,
		&user.LastName,
		&user.Password)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Update(user *User) error {
	query := `
	UPDATE users
	SET name = $1, last_name = $2, password = $3, updated_at = $4
	WHERE id = $6`

	args := []interface{}{user.Name, user.LastName, user.Password, time.Now(), user.ID}

	return m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

}

func (m UserModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	DELETE FROM users
	WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
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

func (m MockUserModel) Update(user *User) error {
	return nil
}

func (m MockUserModel) Delete(id int64) error {
	return nil
}
