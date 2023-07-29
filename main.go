package main

import "fmt"

func main() {
	db, err := NewPostgresDb()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", db)
	if err := db.Start(); err != nil {
		panic(err)
	}
	server := NewServer(":3001", db)
	server.Start()
	fmt.Println("Starting server")
}
