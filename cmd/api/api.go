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

func intR3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("route 3"))
}

func intR4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("route 4"))
}

func (s *APIServer) Run() error {
	extR := http.NewServeMux()
	extR.HandleFunc("GET /1", extR1)
	extR.HandleFunc("GET /2", extR2)

	intR := http.NewServeMux()
	intR.HandleFunc("GET /3", intR3)
	intR.HandleFunc("GET /4", intR4)

	rootR := http.NewServeMux()
	rootR.Handle("/ext/", http.StripPrefix("/ext", extR))
	rootR.Handle("/int/", http.StripPrefix("/int", intR))

	server := http.Server{
		Addr:    s.baseURL,
		Handler: rootR,
	}
	return server.ListenAndServe()
}
