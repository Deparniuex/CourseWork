package httpserver

import "time"

type Option func(*Server)

func WithPort(port string) Option {
	return func(s *Server) {
		s.server.Addr = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}
