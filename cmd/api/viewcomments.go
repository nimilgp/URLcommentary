package main

import (
	"fmt"
	"net/http"
)

func (app *application) getNewestComments(w http.ResponseWriter, r *http.Request) {
	js := `{
		"new_comments": [
			{
				"user_id": "mathew",
				"comment_content": "In another life, I would have really liked just doing laundry and taxes with you." 
			},
			{
				"user_id" : "mark",
				"comment_content": "I am Iron Man." 
			},
			{
				"user_id": "luke",
				"comment_content": "You are a sad, strange, little man, and you have my pity. Farewell." 
			},
			{
				"user_id": "john",
				"comment_content":"Never gonna give you up, Never gonna let you down"
			},
			{
				"user_id": "doe",
				"comment_content": "hello how do you do?" 
			}
		]
	}`

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}

func (app *application) getOldestComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getOldestComments")
}

func (app *application) getRelevantComments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getRelevantComments")
}

func (app *application) getReplyComments(w http.ResponseWriter, r *http.Request) {
	js := `{
		"new_comments": [
			{
				"user_id": "john",
				"comment_content":"Never gonna give you up, Never gonna let you down"
			},
			{
				"user_id": "doe",
				"comment_content": "hello how do you do?" 
			}
		]
	}`

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
