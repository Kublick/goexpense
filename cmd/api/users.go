package main

import (
	"fmt"
	"net/http"
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
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
