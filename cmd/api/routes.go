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
	subR.HandleFunc("GET /page/details", s.getPagedetails)
	subR.HandleFunc("GET /newest/comments/{pageid}/{offset}", s.getNewestParentComments)
	subR.HandleFunc("GET /oldest/comments/{pageid}/{offset}", s.getOldestParentComments)
	subR.HandleFunc("GET /child/comments/{pageid}/{commentid}/{offset}", s.getChildComments)

	return s.addVersionPrefix(subR)
}
