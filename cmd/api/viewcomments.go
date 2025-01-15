package main

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
)

const (
	LIMIT = 5
)

func (app *application) getPageDetails(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("forurl")
	pageDetails, err := app.queries.RetrievePageDetails(app.ctx, url)
	if err != nil {
		app.logger.Warn("get page details failed")
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"page_details": pageDetails}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (app *application) getNewestComments(w http.ResponseWriter, r *http.Request) {
	pageid, err := app.readIDParam(r, "pageid")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	args := dblayer.RetrieveNewstCommentsParams{
		Pageid: pageid,
		Limit:  LIMIT,
	}
	comments, _ := app.queries.RetrieveNewstComments(app.ctx, args)
	app.writeJSON(w, http.StatusOK, envelope{"newest_comments": comments}, nil)
}

func (app *application) getOldestComments(w http.ResponseWriter, r *http.Request) {
	// start, err := app.readIDParam(r, "start")
	// if err != nil {
	// 	http.NotFound(w, r)
	// 	return
	// }
	pageid, err := app.readIDParam(r, "pageid")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "start val : %d\n", start)
	// fmt.Fprintf(w, "url val: %s\n", queryValues.Get("forurl"))
	args := dblayer.RetrieveOldestCommentsParams{
		Pageid: pageid,
		Limit:  LIMIT,
	}
	comments, _ := app.queries.RetrieveOldestComments(app.ctx, args)
	app.writeJSON(w, http.StatusOK, envelope{"oldest_comments": comments}, nil)
}

// func (app *application) getRelevantComments(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "getRelevantComments")
// }

func (app *application) getReplyComments(w http.ResponseWriter, r *http.Request) {
	parentid, err := app.readIDParam(r, "commentid")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	args := dblayer.RetrieveSubCommentsParams{
		Parentid: parentid,
		Limit:    LIMIT,
	}
	comments, _ := app.queries.RetrieveSubComments(app.ctx, args)
	app.writeJSON(w, http.StatusOK, envelope{"sub_comments": comments}, nil)
}
