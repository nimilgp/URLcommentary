package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/config"
	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) getHealthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":          "available",
		"deployment-type": "development",
		"version":         config.Cfg.Version,
	}
	err := response.JSON(w, http.StatusOK, envelope{"healthcheck": data})
	if err != nil {
		s.serverError(w, r, err)
	}
}
