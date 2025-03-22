package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) logServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	s.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *APIServer) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		app.logServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *APIServer) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	app.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *APIServer) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *APIServer) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *APIServer) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

// func (app *APIServer) failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
// 	err := response.JSON(w, http.StatusUnprocessableEntity, v)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 	}
// }

func (app *APIServer) invalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", "Bearer")

	app.errorMessage(w, r, http.StatusUnauthorized, "Invalid authentication token", headers)
}

func (app *APIServer) authenticationRequired(w http.ResponseWriter, r *http.Request) {
	app.errorMessage(w, r, http.StatusUnauthorized, "You must be authenticated to access this resource", nil)
}

func (app *APIServer) basicAuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	message := "You must be authenticated to access this resource"
	app.errorMessage(w, r, http.StatusUnauthorized, message, headers)
}
