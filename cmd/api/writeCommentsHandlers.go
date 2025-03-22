package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/request"
	"github.com/nimilgp/URLcommentary/internal/response"
	"github.com/nimilgp/URLcommentary/internal/validator"
)

func (s *APIServer) postParentComment(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pageid    int32               `json: pageid`
		Userid    int32               `json: userid`
		Content   string              `json: content`
		Validator validator.Validator `json: -`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		s.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Pageid != 0, "PageId", "PageId required")
	input.Validator.CheckField(input.Userid != 0, "UserId", "UserId required")
	input.Validator.CheckField(input.Content != "", "Content", "Content required")

	if input.Validator.HasErrors() {
		s.failedValidation(w, r, input.Validator)
		return
	}

	arg := dblayer.CreateParentCommentParams{
		Pageid:  input.Pageid,
		Userid:  input.Userid,
		Content: input.Content,
	}
	commentid, err := s.querier.CreateParentComment(s.ctx, arg)
	if err != nil {
		s.serverError(w, r, err)
	}
	err = response.JSON(w, http.StatusOK, envelope{"comment-id": commentid})
	if err != nil {
		s.serverError(w, r, err)
	}
}
