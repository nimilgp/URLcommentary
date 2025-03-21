package api

import (
	"fmt"
	"net/http"
)

func (s *APIServer) status(w http.ResponseWriter, r *http.Request) {
	email, _ := s.querier.RetrieveUserDetails(s.ctx, "ngp")
	w.Write([]byte(fmt.Sprint(email)))
}
