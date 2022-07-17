package app

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	return &Server{
		mux.NewRouter(),
	}
}

func (s *Server) AddMiddleWares() {
	s.Use(middleWare)
}

func (s *Server) AddRoutes() {
	s.Handle("/", rootHandler())
}
