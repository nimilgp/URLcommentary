package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
)

func (app *application) postComment(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pageid      int32
		Userid      int32
		Commentdata string
		// Parentid    int32
	}
	json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input)
	arg := dblayer.CreateCommentParams{
		Pageid:      input.Pageid,
		Userid:      input.Userid,
		Commentdata: input.Commentdata,
		Parentid:    0,
	}
	app.queries.CreateComment(app.ctx, arg)
}

func (app *application) postReplyComment(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pageid      int32
		Userid      int32
		Commentdata string
		Parentid    int32
	}
	json.NewDecoder(r.Body).Decode(&input)
	arg := dblayer.CreateCommentParams{
		Pageid:      input.Pageid,
		Userid:      input.Userid,
		Commentdata: input.Commentdata,
		Parentid:    input.Parentid,
	}
	app.queries.CreateComment(app.ctx, arg)
}
