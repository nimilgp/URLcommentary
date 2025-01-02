package main

import (
	"fmt"
	"net/http"
)

func (app *application) postComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "postComment")
}

func (app *application) postReplyComment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "postReplyComment")
}
