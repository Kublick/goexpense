package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// Protect routes with middleware of logged user
	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck",  app.requireActivatedUser(app.healthcheckHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.requirePermission("expenses:read", app.getUserById))
	router.HandlerFunc(http.MethodPut, "/v1/users/:id", app.updateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", app.deleteUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/:id/activate", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
