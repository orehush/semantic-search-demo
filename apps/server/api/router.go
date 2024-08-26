package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/hello", Hello)
	return router
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}
