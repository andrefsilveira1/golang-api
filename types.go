package main

import (
	"math/rand"
	"time"
)

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	BankNumber int64     `json:"bank"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PostAccount struct {
	Name string `json:"name"`
}

func NewAccount(name string) *Account {
	return &Account{
		Name:       name,
		BankNumber: int64(rand.Intn(100000)),
		CreatedAt:  time.Now(),
	}
}

type transferRequest struct {
	ToAccount int       `json:"toAccount"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date"`
}
