package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)

	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Hello just sends a hello world message back to the user
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
