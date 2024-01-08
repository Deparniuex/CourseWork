package httpserver

import (
	"net/http"
)

type Server struct {
	server *http.Server
	notify chan error
}

func NewServer(handler http.Handler, opts ...Option) *Server {
	httpserver := &http.Server{
		Handler: handler,
	}
	s := &Server{
		server: httpserver,
		notify: make(chan error, 1),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
