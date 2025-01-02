package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.getHealthcheck)
	router.HandlerFunc(http.MethodGet, "/view/newest/comments/:start/:stop", app.getNewestComments)
	router.HandlerFunc(http.MethodGet, "/view/oldest/comments/:start/:stop", app.getOldestComments)
	router.HandlerFunc(http.MethodGet, "/view/relevant/comments/:start/:stop", app.getRelevantComments)
	router.HandlerFunc(http.MethodPost, "/create/comment", app.postComment)
	router.HandlerFunc(http.MethodGet, "/view/reply/comments/", app.getReplyComments)
	router.HandlerFunc(http.MethodPost, "/create/reply/comment", app.postReplyComment)

	return router
}
