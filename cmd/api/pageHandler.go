package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) getPagedetails(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("forurl")
	pageDetails, err := s.querier.RetrievePageDetails(s.ctx, url)
	if err != nil {
		s.logger.Warn("get page details failed")
	}
	err = response.JSON(w, http.StatusOK, envelope{"healthcheck": pageDetails})
	if err != nil {
		s.serverError(w, r, err)
	}
}
