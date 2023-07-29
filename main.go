package main

import "fmt"

func main() {
	db, err := NewPostgresDb()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", db)
	server := NewServer(":3001", db)
	server.Start()
	fmt.Println("Starting server")
}
