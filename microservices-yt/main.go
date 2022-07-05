package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("HEY THERE!"))
		log.Println("Hello World!")
	})

	// bind to every IP (:9090); constructs an HTTP server and registers a default handler
	// default serve mux
	http.ListenAndServe(":9090", nil)
}
