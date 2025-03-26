package api

import (
	"net/http"
	"strconv"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/request"
	"github.com/nimilgp/URLcommentary/internal/response"
	"github.com/nimilgp/URLcommentary/internal/validator"
)

func (s *APIServer) getUserDetails(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("userid")
	userId, err := strconv.Atoi(param)
	if err != nil {
		s.logger.Info("invalid param", "userid", param)
	}

	details, err := s.querier.RetrieveUserDetails(s.ctx, int32(userId))
	if err != nil {
		s.logger.Warn("get user details failed")
	}
	err = response.JSON(w, http.StatusOK, envelope{"user-details": details})
	if err != nil {
		s.serverError(w, r, err)
	}
}

func (s *APIServer) putUserDetails(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Userid    int32               `json: Userid`
		Username  string              `json: Username`
		Fullname  string              `json: Fullname`
		Aboutme   string              `json: Aboutme`
		Validator validator.Validator `json: -`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		s.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Userid != 0, "UserId", "Userid required")
	input.Validator.CheckField(input.Username != "", "Username", "Username required")
	input.Validator.CheckField(input.Fullname != "", "Fullname", "Fullname required")
	input.Validator.CheckField(input.Aboutme != "", "Aboutme", "Aboutme required")

	if input.Validator.HasErrors() {
		s.failedValidation(w, r, input.Validator)
		return
	}

	arg := dblayer.UpdateUserDetailsParams{
		Username: input.Username,
		Fullname: input.Fullname,
		Aboutme:  input.Aboutme,
		Userid:   input.Userid,
	}

	err = s.querier.UpdateUserDetails(s.ctx, arg)
	if err != nil {
		s.serverError(w, r, err)
	}
	err = response.JSON(w, http.StatusOK, envelope{"user-details": "updated"})
	if err != nil {
		s.serverError(w, r, err)
	}
}
