package utils

import (
	"fmt"
	"net/http"
)

func Version() string {
	return "v0.0.1"
}

func NewServer() *http.Server {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	err := server.ListenAndServe()
	if err != nil {
		return nil
	}
	return &server
}
