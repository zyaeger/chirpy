package main

import (
	"log"
	"net/http"
)

func main() {
	smu := http.NewServeMux()
	server := http.Server{
		Handler: smu,
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("issue starting server: %v\n", err)
		return
	}
}