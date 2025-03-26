package api

import (
	"net/http"
	"strconv"

	"github.com/nimilgp/URLcommentary/internal/response"
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
