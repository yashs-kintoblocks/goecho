package main

// IntroMessage is the payload that gets sent back on GET /
type IntroMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
