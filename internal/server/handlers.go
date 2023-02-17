package server

import (
	"bankAPI/internal/model"
	"encoding/json"
	"net/http"
)

// Handlers configuration setup
func (s *Server) ConfigureRouter() {
	// Root directory
	s.Router.HandleFunc("/", s.RootGet).Methods("GET")
	s.Router.HandleFunc("/", s.RootPost).Methods("POST")
	// Registration and authentication
	s.Router.HandleFunc("/sign_up", s.SignUp).Methods("POST")
	s.Router.HandleFunc("/sign_in", s.SignIn).Methods("POST")
}

// Handles GET method in root directory
func (s *Server) RootGet(w http.ResponseWriter, r *http.Request) {

}

// Handles POST method in root directory
func (s *Server) RootPost(w http.ResponseWriter, r *http.Request) {

}

// Handles registartion
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	type Respond struct {
		Message string
	}
	// Parsing request
	req := &Request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	// Validating data
	// ...
	// Checking user in database
	u := new(model.User)
	u.Auth.Login = req.Login
	u.Auth.Password = req.Password
	exist, err := s.Storage.CheckUser(u)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	if exist {
		s.Respond(w, r, http.StatusConflict, "This login is already taken")
		return
	}
	// Creating user
	err = s.Storage.CreateUser(u)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	s.Respond(w, r, http.StatusCreated, "User created successfully")
}

// Handles authentication
func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	type Respond struct {
		Message string
	}
	// Parsing request
	req := &Request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	// Checking user in database
	u := new(model.User)
	u.Auth.Login = req.Login
	u.Auth.Password = req.Password
	exist, err := s.Storage.CheckUser(u)
	if err != nil {
		s.Error(w, r, http.StatusBadRequest, err)
		return
	}
	if !exist {
		s.Respond(w, r, http.StatusUnauthorized, "User doesn't exist")
		return
	}
	// Verifying password
	verified, err := s.Storage.VerifyPassword(u)
	if !verified {
		s.Respond(w, r, http.StatusUnauthorized, "Wrong password")
		return
	}
	s.Respond(w, r, http.StatusAccepted, "Access")
}
