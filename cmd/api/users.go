package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kublick/goexpense/internal/data"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "register a new user")
}

func (app *application) getUserById(w http.ResponseWriter, r *http.Request) {
	// TODO might change to read a string insterad of number id
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	user := data.User{
		ID:        id,
		Name:      "test",
		LastName:  "test",
		Email:     "test@test.com",
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
