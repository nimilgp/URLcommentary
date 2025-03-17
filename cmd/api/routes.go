package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /status", app.status)
	mux.HandleFunc("POST /users", app.createUser)
	mux.HandleFunc("POST /authentication-tokens", app.createAuthenticationToken)

	mux.Handle("GET /protected", app.requireAuthenticatedUser(http.HandlerFunc(app.protected)))

	mux.Handle("GET /basic-auth-protected", app.requireBasicAuthentication(http.HandlerFunc(app.protected)))

	return app.logAccess(app.recoverPanic(app.authenticate(mux)))
}
