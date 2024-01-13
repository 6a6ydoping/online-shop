package httpserver

import "time"

type Option func(*Server)

func WithReadTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.server.WriteTimeout = timeout
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.shutdownTimeout = timeout
	}
}
