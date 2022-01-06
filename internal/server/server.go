package server

import (
	v1 "calendar/internal/server/v1"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewServer() *Server{
	return &Server{
		router: v1.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *Server) Run() error{
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
