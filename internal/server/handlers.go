package server

import (
	"net/http"
)

// Handlers configuration setup
func (s *Server) ConfigureRouter() {
	// Root directory
	s.Router.HandleFunc("/", s.RootGetHandler).Methods("GET")
	s.Router.HandleFunc("/", s.RootPostHandler).Methods("POST")
	// Auth directory
	s.Router.HandleFunc("/auth", s.RootGetHandler).Methods("GET")
	s.Router.HandleFunc("/auth", s.RootPostHandler).Methods("POST")
}

// Handles GET method in root directory
func (s *Server) RootGetHandler(w http.ResponseWriter, r *http.Request) {

}

// Handles POST method in root directory
func (s *Server) RootPostHandler(w http.ResponseWriter, r *http.Request) {

}

// Handles GET method in /auth
func (s *Server) AuthGetHandler(w http.ResponseWriter, r *http.Request) {

}

// Handles POST method in /auth
func (s *Server) AuthPostHandler(w http.ResponseWriter, r *http.Request) {

}
