package server

import (
	"net/http"
)

const stackEmpty = "stack is empty"
const missingValue = "missing value"

func (s *Server) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.Handle("/push", http.HandlerFunc(s.HandlePush))
	mux.Handle("/top", http.HandlerFunc(s.HandleTop))
	mux.Handle("/pop", http.HandlerFunc(s.HandlePop))

	return mux
}

func (s *Server) HandlePush(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	value := request.FormValue("value")
	if value == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(missingValue))

		return
	}

	s.stack.Push(value)
	writer.WriteHeader(http.StatusOK)
}

func (s *Server) HandlePop(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	value, ok := s.stack.Pop()
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte(stackEmpty))

		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(value))
}

func (s *Server) HandleTop(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	value, ok := s.stack.Top()
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte(stackEmpty))

		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte(value))
}
