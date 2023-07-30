package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccount(int) (*Account, error)
	GetAccounts() ([]*Account, error)
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

func (s *PostStore) Start() error {
	return s.CreateAccountTable()
}

func (s *PostStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account (name, number, balance, created_at) values ($1, $2, $3, $4)`
	res, err := s.db.Query(query, acc.Name, acc.BankNumber, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}
	fmt.Printf("%+v \n", res)

	return nil
}

func (s *PostStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id = $1", id)
	return err
}
func (s *PostStore) GetAccount(id int) (*Account, error) {
	query, err := s.db.Query("SELECT * FROM account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for query.Next() {
		return searchAccount(query)
	}
	return nil, nil
}

func (s *PostStore) GetAccounts() ([]*Account, error) {
	res, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	accs := []*Account{}
	for res.Next() {
		account, err := searchAccount(res)
		if err != nil {
			return nil, err
		}
		accs = append(accs, account)
	}

	return accs, nil

}

func searchAccount(res *sql.Rows) (*Account, error) {
	account := Account{}
	if err := res.Scan(&account.Id, &account.Name, &account.BankNumber, &account.Balance, &account.CreatedAt); err != nil {
		return nil, err
	}
	return &account, nil
}
