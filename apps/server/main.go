package main

import (
	"log"
	"net/http"
	"semantic-search-demo/internal/api"
)

func main() {
	router := api.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
