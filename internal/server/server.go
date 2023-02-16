package server

import (
	"bankAPI/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router *mux.Router
	Logger *logrus.Logger
	Store  *storage.Storage
}
