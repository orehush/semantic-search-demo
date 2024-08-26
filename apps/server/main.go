package main

import (
	"log"
	"net/http"
	"os"

	"semantic-search-demo/src/api"
	"semantic-search-demo/src/app"

	"github.com/julienschmidt/httprouter"
)

func main() {
	app.InitDB()

	router := httprouter.New()
	router.POST("/admin/synonyms", api.AddSynonymsHandler)
	router.GET("/synonyms", api.GetSynonymsHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
