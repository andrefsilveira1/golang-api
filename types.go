package main

import "math/rand"

type Account struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	BankNumber int64  `json:"bank"`
	Balance    int64  `json:"balance"`
}

func NewAccount(name string) *Account {
	return &Account{
		Id:         rand.Intn(100000),
		Name:       name,
		BankNumber: int64(rand.Intn(100000000000000)),
	}
}
