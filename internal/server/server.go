package server

import (
	"net/http"
	"os"
	"os/signal"
	"stack-service/internal/config"
	"stack-service/internal/stack"
	"syscall"
	"time"
)

const connectTimeOut = 3 * time.Second

type Server struct {
	defaultPort string
	stack       *stack.Stack
}

func New(cfg *config.Config, stack *stack.Stack) *Server {
	return &Server{
		defaultPort: cfg.DefaultPort,
		stack:       stack,
	}
}

func (s *Server) Run() {
	go func() {
		//nolint:exhaustivestruct,exhaustruct
		server := &http.Server{
			Addr:              s.defaultPort,
			ReadHeaderTimeout: connectTimeOut,
			Handler:           s.InitRoutes(),
		}

		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal
}
