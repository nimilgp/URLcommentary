package main

import (
	"fmt"
	"net/http"
)

func (app *application) getNewestComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getNewestComments")
}

func (app *application) getOldestComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getOldestComments")
}

func (app *application) getRelevantComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getRelevantComments")
}

func (app *application) getReplyComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getReplyComments")
}
