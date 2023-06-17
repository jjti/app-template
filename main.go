package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("starting server")

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("got a hello request")
		fmt.Fprintf(w, "hello\n")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(fmt.Printf("failed to host service: %v", err))
	}
}
