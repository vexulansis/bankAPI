package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Respond(w http.ResponseWriter, r *http.Request, code int, data any) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
func (s *Server) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.Respond(w, r, code, map[string]string{"error": err.Error()})
}
