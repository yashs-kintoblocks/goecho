package main

import (
	"encoding/json"
	"net/http"
)

// Intro just sends a Intro world message back to the user
func Intro(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(RootResponse{
		Status:  "ok",
		Message: "GET /get with query params or POST /post with a body to hear and echo",
	})

	if err != nil {
		panic(err)
	}
}

// GetEcho just responds with the query parameters given to the call
func GetEcho(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	jsonEnc := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := jsonEnc.Encode(GetEchoResponse{Status: "ok", QueryParams: queryParams})

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		jsonEnc.Encode(ErrorResponse{Status: "error", Message: err})
	}
}

// PostEcho just responds with the query parameters given to the call
func PostEcho(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(RootResponse{
		Status:  "ok",
		Message: "This is where the POST body will go",
	})

	if err != nil {
		panic(err)
	}
}
