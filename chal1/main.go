package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":12345", http.FileServer(http.Dir("/go/src/chal1"))))
}
