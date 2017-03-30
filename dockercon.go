package main

import (
	"fmt"
	"log"
	"net/http"
)

const Version string = "1.0"

// version 1.0
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got a request!")
		fmt.Fprintf(w, "Version %v!\n", Version)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
