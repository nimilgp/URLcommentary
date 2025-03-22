package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) getPagedetails(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("forurl")

	exist, err := s.querier.DoesPageExist(s.ctx, url)
	if err != nil {
		s.logger.Warn("page existance could not be verified")
	}
	if !exist {
		s.querier.CreatePage(s.ctx, url)
		s.logger.Info("page created", "url", url)
	}

	pageDetails, err := s.querier.RetrievePageDetails(s.ctx, url)
	if err != nil {
		s.logger.Warn("get page details failed")
	}
	err = response.JSON(w, http.StatusOK, envelope{"page-details": pageDetails})
	if err != nil {
		s.serverError(w, r, err)
	}
}
