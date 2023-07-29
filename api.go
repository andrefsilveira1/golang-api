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
	db      Database
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func NewServer(addrs string, database Database) *Server {
	return &Server{
		address: addrs,
		db:      database,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))
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
	id := mux.Vars(r)["id"]
	fmt.Println("ID:", id)
	return WriteJson(w, http.StatusOK, &Account{})
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
