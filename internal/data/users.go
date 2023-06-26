package data

import "time"

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updatedAt"`
}
