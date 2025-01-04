package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.getHealthcheck)
	router.HandlerFunc(http.MethodGet, "/view/newest/comments/:forurl/:start/", app.getNewestComments)
	router.HandlerFunc(http.MethodGet, "/view/oldest/comments/:forurl/:start/", app.getOldestComments)
	router.HandlerFunc(http.MethodGet, "/view/relevant/comments/:forurl/:start/", app.getRelevantComments)
	router.HandlerFunc(http.MethodPost, "/create/comment/:forurl", app.postComment)
	router.HandlerFunc(http.MethodGet, "/view/reply/comments/:forurl/:commentid/:start/", app.getReplyComments)
	router.HandlerFunc(http.MethodPost, "/create/reply/comment/:forurl/:commentid", app.postReplyComment)

	return app.enableCORS(router)
}
