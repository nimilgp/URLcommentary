package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/request"
	"github.com/nimilgp/URLcommentary/internal/validator"
)

func (s *APIServer) postLike(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pageid    int32               `json: Pageid`
		Userid    int32               `json: Userid`
		Commentid int32               `json: Commentid`
		Likevalue int32               `json: Likevalue`
		Validator validator.Validator `json: -`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		s.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Pageid != 0, "PageId", "Pageid required")
	input.Validator.CheckField(input.Userid != 0, "UserId", "Userid required")
	input.Validator.CheckField(input.Commentid != 0, "Commentid", "Commentid required")
	input.Validator.CheckField(input.Likevalue != 0, "Commentid", "Likevalue required(-1/+1)")

	if input.Validator.HasErrors() {
		s.failedValidation(w, r, input.Validator)
		return
	}

	arg1 := dblayer.CreateLikeHistoryParams{
		Pageid:    input.Pageid,
		Userid:    input.Userid,
		Commentid: input.Commentid,
		Likevalue: input.Likevalue,
	}

	arg2 := dblayer.DoesLikeExistParams{
		Pageid:    input.Pageid,
		Userid:    input.Userid,
		Commentid: input.Commentid,
	}

	arg3 := dblayer.UpdateLikeHistoryParams{
		Pageid:    input.Pageid,
		Userid:    input.Userid,
		Commentid: input.Commentid,
		Likevalue: input.Likevalue,
	}

	exist, err := s.querier.DoesLikeExist(s.ctx, arg2)
	if err != nil {
		s.logger.Warn("like existance could not be verified")
		return
	}

	if exist {
		err := s.querier.UpdateLikeHistory(s.ctx, arg3)
		if err != nil {
			s.serverError(w, r, err)
			return
		}
	} else {
		err = s.querier.CreateLikeHistory(s.ctx, arg1)
		if err != nil {
			s.serverError(w, r, err)
			return
		}
	}
}
