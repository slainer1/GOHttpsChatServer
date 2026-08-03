package httpServer

import "net/http"

// route the users request
//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		switch r.URL.Path {
//		case "/":
//			w.Write([]byte("index page"))
//			return
//		case "/users":
//			w.Write([]byte("users page"))
//			return
//		}
//	default:
//		w.Write([]byte("404 page not found"))
//		return
//	}
//}

func main() {
	api := &api{addr: ":8080"}

	//initalize the mux for pausing
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
