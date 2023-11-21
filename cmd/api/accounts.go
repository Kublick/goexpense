package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kublick/goexpense/internal/data"
	"github.com/kublick/goexpense/internal/validator"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name        string  `json:"name"`
		AccountType string  `json:"type"`
		Balance     float64 `json:"balance"`
		Budget      bool    `json:"budget"`
		UserID      int64   `json:"userId"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Println(input)
	account := &data.Account{
		Name:        input.Name,
		AccountType: input.AccountType,
		Balance:     input.Balance,
		Budget:      input.Budget,
		UserID:      input.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	v := validator.New()

	if data.ValidateAccount(v, account); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Accounts.Insert(account)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	account, err := app.models.Accounts.GetAll(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
