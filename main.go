package main

import (
	"fmt"
	"log"
)

// TODO

func main() {
	store, err := NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	
	server := NewAPIServer(":3000")
	server.Run()
}
