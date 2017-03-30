package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Version        string = "2.1.1"
	GordonGreeting string = "I sure love Dockercon!"
)

// version 1.0
func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		// TODO(dperny) don't let any miscreants deploy without gordon's greeting!
		if len(GordonGreeting) == 0 {
			// TODO(NOT SCRIPT)(dperny) write an error handler in v2.0 that logs
			// errors before returning them
			http.Error(w, "forgot to give gordon a greeting!", http.StatusInternalServerError)
			return
		}
		log.Printf("Got a healthcheck!")
		fmt.Fprintln(w, "Healthy!")
	})

	http.HandleFunc("/gordon", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got a request for gordon!")
		fmt.Fprintf(w, "gordon says %v!\n", GordonGreeting)
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
