package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"C:\Users\ReissT\.vscode\working\.vscode\working\handlers"
)

func main() {
	// Hello world, the web server
/*
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)

	goodbyeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Bye World!\n")
	}
	http.HandleFunc("/bye", goodbyeHandler)
	*/

	l := log.New(os.Stdout, "")
	hh := handlers.newHello()
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
