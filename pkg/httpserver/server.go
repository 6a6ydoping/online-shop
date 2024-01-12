package httpserver

import (
	"net/http"
	"time"
)

type Server struct {
	server       *http.Server
	shutdownTime time.Duration
	notify       chan error
}

func New(h http.Handler, opts ...Option) *Server {
	return nil
}
