package api

import "net/http"

type APIServer struct {
	baseURL string
}

func GetAPIServer(baseURL string) *APIServer {
	return &APIServer{
		baseURL: baseURL,
	}
}

func extR1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("route 1"))
}

func extR2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("route 2"))
}

func (s *APIServer) Run() error {
	subR := http.NewServeMux()
	subR.HandleFunc("GET /1", extR1)
	subR.HandleFunc("GET /2", extR2)

	rootR := http.NewServeMux()
	rootR.Handle("/api/v1/", http.StripPrefix("/api/v1", subR))
	server := http.Server{
		Addr:    s.baseURL,
		Handler: rootR,
	}
	return server.ListenAndServe()
}
