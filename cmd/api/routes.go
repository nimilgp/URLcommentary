package api

import "net/http"

func (s *APIServer) addVersionPrefix(subR *http.ServeMux) http.Handler {
	rootR := http.NewServeMux()
	prefix := "/api/v" + s.version + "/"
	rootR.Handle(prefix, http.StripPrefix("/api/v"+s.version, subR))
	return rootR
}

func (s *APIServer) routes() http.Handler {
	subR := http.NewServeMux()
	subR.HandleFunc("GET /healthcheck", s.getHealthcheck)

	return s.addVersionPrefix(subR)
}
