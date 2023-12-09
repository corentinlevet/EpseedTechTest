package server

import (
	"epseed/internal/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func logCallMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Appel de l'API:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/signup", handler.SignupHandler).Methods("POST")
	return r
}

func NotesRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/notes/get", handler.GetNotesForUserHandler).Methods("GET")
	r.HandleFunc("/notes/create", handler.CreateNoteHandler).Methods("POST")
	r.HandleFunc("/notes/update", handler.UpdateNoteHandler).Methods("PUT")
	r.HandleFunc("/notes/delete", handler.DeleteNoteHandler).Methods("DELETE")
	return r
}

func UserRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/get", handler.GetUsersHandler).Methods("GET")
	return r
}

func InitRoutes() {
	r := mux.NewRouter()
	r.Use(logCallMiddleware)

	r.PathPrefix("/auth").Handler(AuthRouter())
	r.PathPrefix("/notes").Handler(NotesRouter())
	r.PathPrefix("/users").Handler(UserRouter())

	http.Handle("/", r)
}
