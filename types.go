package main

import "math/rand"

type Account struct {
	Id         int
	Name       string
	BankNumber int64
	Balance    int64
}

func NewAccount(name string) *Account {
	return &Account{
		Id:         rand.Intn(100000),
		Name:       name,
		BankNumber: int64(rand.Intn(100000000000000)),
	}
}
