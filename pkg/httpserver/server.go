package httpserver

import (
	"github.com/valyala/fasthttp"
	"time"
)

const serverPort = ":8080"

type Server struct {
	server          *fasthttp.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(h *fasthttp.RequestHandler, opts ...Option) *Server {
	httpServer := &fasthttp.Server{
		Handler: *h,
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe(serverPort)
		close(s.notify)
	}()
}

func (s *Server) Shutdown() error {
	//ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	//
	//defer cancel()

	return s.server.Shutdown()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
