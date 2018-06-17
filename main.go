package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Intro)
	r.HandleFunc("/get", GetEcho).Methods("GET")
	r.HandleFunc("/post", PostEcho).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Starting server at 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
