package server

import (
	"bankAPI/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router  *mux.Router
	Logger  *logrus.Logger
	Storage *storage.Storage
}

// Server factory
func NewServer(storage *storage.Storage) *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		Logger:  logrus.New(),
		Storage: storage,
	}
	return s
}

// Starting method
func (s *Server) Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}
	err = http.ListenAndServe(":"+os.Getenv("SERVERPORT"), handlers.CORS()(s.Router))
	if err != nil {
		log.Fatal()
	}
}
