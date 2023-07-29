package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccount(int) (*Account, error)
}

type PostStore struct {
	db *sql.DB
}

func NewPostgresDb() (*PostStore, error) {
	connectionStr := "user=postgres dbname=postgres password=root sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostStore{
		db: db,
	}, nil
}

func (s *PostStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostStore) GetAccount(id int) (*Account, error) {
	return nil, nil
}
