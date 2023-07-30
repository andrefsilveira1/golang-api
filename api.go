package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleGetAccounts))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))
	fmt.Println("API server running on port: ", s.address)
	http.ListenAndServe(s.address, router)
}
func (s *Server) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateAccount(w, r)
	default:
		WriteJson(w, http.StatusBadRequest, apiError{Error: "Invalid Method"})

	}
	return nil
}

func (s *Server) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accs, err := s.db.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, accs)
}

func (s *Server) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := convertId(r)
		if err != nil {
			return err
		}

		account, err := s.db.GetAccount(id)
		if err != nil {
			return err
		}

		return WriteJson(w, http.StatusOK, account)
	}
	if r.Method == "DELETE" {
		return s.handleDelete(w, r)
	}
	return nil
}

func (s *Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := PostAccount{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc := NewAccount(req.Name)
	if err := s.db.CreateAccount(acc); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, acc)
}

func (s *Server) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	req := &transferRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	defer r.Body.Close()
	return WriteJson(w, http.StatusOK, r.Body)
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) error {
	id, err := convertId(r)
	if err != nil {
		return err
	}
	if err := s.db.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, map[string]int{"deleted": id})
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

func convertId(r *http.Request) (int, error) {
	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return id, err
	}

	return id, nil
}
