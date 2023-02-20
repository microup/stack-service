package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"stack-service/internal/config"
	"stack-service/internal/stack"
	"syscall"
)

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
		if err := http.ListenAndServe(s.defaultPort, s.InitRoutes()); err != nil {
			log.Panic(err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	<-stopSignal
}
