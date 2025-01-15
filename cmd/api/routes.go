package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.getHealthcheck)
	router.HandlerFunc(http.MethodGet, "/view/page/details", app.getPageDetails)
	router.HandlerFunc(http.MethodGet, "/view/newest/comments/:pageid/:start", app.getNewestComments)
	router.HandlerFunc(http.MethodGet, "/view/oldest/comments/:pageid/:start", app.getOldestComments)
	// router.HandlerFunc(http.MethodGet, "/view/relevant/comments/:start", app.getRelevantComments)
	router.HandlerFunc(http.MethodPost, "/create/comment", app.postComment)
	router.HandlerFunc(http.MethodGet, "/view/reply/comments/:commentid/:start", app.getReplyComments)
	router.HandlerFunc(http.MethodPost, "/create/reply/comment/:commentid/:start", app.postReplyComment)

	return app.enableCORS(router)
}
