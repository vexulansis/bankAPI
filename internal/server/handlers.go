package server

import (
	"bankAPI/internal/model"
	"encoding/json"
	"net/http"
)

// Handlers configuration setup
func (s *Server) ConfigureRouter() {
	// Root directory
	s.Router.HandleFunc("/", s.RootGetHandler).Methods("GET")
	s.Router.HandleFunc("/", s.RootPostHandler).Methods("POST")
	// Auth directory
	s.Router.HandleFunc("/auth", s.AuthGetHandler).Methods("GET")
	s.Router.HandleFunc("/auth", s.AuthPostHandler).Methods("POST")
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
	type Request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	req := &Request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	u := new(model.User)
	u.Auth.Login = req.Login
	u.Auth.Password = req.Password
	err = s.Storage.CreateUser(u)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
	}
	s.Respond(w, r, http.StatusCreated, u)
}
