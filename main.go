package main

import "fmt"

func main() {
	server := NewServer(":3001")
	server.Start()
	fmt.Println("Starting server")
}
