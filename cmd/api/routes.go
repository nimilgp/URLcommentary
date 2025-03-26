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
	subR.HandleFunc("POST /parent/comment", s.postParentComment)
	subR.HandleFunc("POST /child/comment", s.postChildComment)
	subR.HandleFunc("POST /like", s.postLike)
	subR.HandleFunc("GET /like/{pageid}/{userid}", s.getLikeHistory)
	subR.HandleFunc("POST /signup", s.postSignUp)
	subR.HandleFunc("POST /signin", s.postSignIn)
	subR.HandleFunc("GET /user/details/{userid}", s.getUserDetails)
	return s.logAccess(s.addVersionPrefix(subR))
}
