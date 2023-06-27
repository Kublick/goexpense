package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kublick/goexpense/internal/data"
	"github.com/kublick/goexpense/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Start validation
	v := validator.New()

	user := data.User{
		Name:     input.Name,
		LastName: input.LastName,
		Email:    input.Email,
		Password: input.Password,
	}

	if data.ValidateUser(v, &user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) getUserById(w http.ResponseWriter, r *http.Request) {
	// TODO might change to read a string insterad of number id
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}

}