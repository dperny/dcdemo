package main

import (
	"fmt"
	"log"
	"net/http"
)

const Version string = "2.0"

// version 1.0
func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got a healthcheck!")
		fmt.Fprintln(w, "Healthy!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound := fmt.Sprintf("%v not found!", r.URL.Path)
			log.Printf("Error! %v", notFound)
			http.Error(w, notFound, http.StatusNotFound)
			return
		}
		log.Printf("Got a request!")
		fmt.Fprintf(w, "Version %v!\n", Version)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
