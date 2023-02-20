package server

import (
	"net/http"
)

func (s *Server) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.Handle("/push", http.HandlerFunc(s.HandlePush))
	mux.Handle("/top", http.HandlerFunc(s.HandleTop))
	mux.Handle("/pop", http.HandlerFunc(s.HandlePop))

	return mux
}

func (s *Server) HandlePush(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value := r.FormValue("value")
	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing value"))
		return
	}

	s.stack.Push(value)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandlePop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := s.stack.Pop()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("stack is empty"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (s *Server) HandleTop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := s.stack.Top()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("stack is empty"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}
