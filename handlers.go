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
		w.WriteHeader(http.StatusInternalServerError)
		jsonEnc.Encode(ErrorResponse{Status: "error", Message: err})
	}
}

// PostEcho just responds with the query parameters given to the call
func PostEcho(w http.ResponseWriter, req *http.Request) {
	jsonEnc := json.NewEncoder(w)
	jsonDec := json.NewDecoder(req.Body)

	var b map[string]string

	decErr := jsonDec.Decode(&b)

	if decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		encErr := json.NewEncoder(w).Encode(PostEchoResponse{Status: "ok", Body: b})

		if encErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonEnc.Encode(ErrorResponse{Status: "error", Message: encErr})
		}
	}
}
