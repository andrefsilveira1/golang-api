package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type apiError struct {
	Error string
}
type Server struct {
	address string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})

		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewServer(addrs string) *Server {
	return &Server{
		address: addrs,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	fmt.Println("API server running on port: ", s.address)
	http.ListenAndServe(s.address, router)
}
func (s *Server) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDelete(w, r)
	default:
		WriteJson(w, http.StatusBadRequest, apiError{Error: "Invalid Method"})

	}
	return nil
}

func (s *Server) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	acc := NewAccount("André")
	return WriteJson(w, http.StatusOK, acc)
}

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
