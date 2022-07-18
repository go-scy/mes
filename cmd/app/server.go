package app

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	csrfKey = "csrf-auth-key"
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

func (s *Server) AddCsrfMiddleware() {
	csrfMiddleWare := csrf.Protect([]byte(csrfKey))
	s.Use(csrfMiddleWare)
}

func (s *Server) AddRoutes() {
	s.Handle("/", rootHandler())
	s.Handle("/employee", employeeHandler())
}
