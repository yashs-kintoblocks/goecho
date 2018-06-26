package main

import (
	"net/url"
)

// RootResponse is the payload that gets sent back on GET /
type RootResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GetEchoResponse is the payload to send back on GET /get
type GetEchoResponse struct {
	Status      string     `json:"status"`
	QueryParams url.Values `json:"queryParams"`
}

// PostEchoResponse is the payload to send back on POST /post
type PostEchoResponse struct {
	Status string            `json:"status"`
	Body   map[string]string `json:"body"`
}

// ErrorResponse is what gets sent when there is an error
type ErrorResponse struct {
	Status  string `json:"status"`
	Message error  `json:"message"`
}
