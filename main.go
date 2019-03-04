package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting on port 8080")

	http.HandleFunc("/", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
