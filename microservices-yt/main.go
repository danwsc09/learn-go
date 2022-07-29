package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("HEY THERE!"))
		log.Println("Goodbye World!")
	})

	// bind to every IP (:9090); constructs an HTTP server and registers a default handler
	// default serve mux
	http.ListenAndServe(":9090", nil)
}
