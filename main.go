package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	s := NewServer()
	http.ListenAndServe("0.0.0.0:9999", s)

	db, _ := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")

	fmt.Println(db.Ping())
}

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
   s.router.HandleFunc("/hi", s.handleHelloWorld).Methods("GET")
}

func (s *Server) handleHelloWorld(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(200)
   msg := "hellow world"
   w.Write([]byte(msg))
}
