package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	"github.com/hemuku90/restAPI/handler"
	"github.com/hemuku90/restAPI/middleware"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", middleware.Chain(handler.GetBooks, middleware.Method("GET"), middleware.Logging()))
	log.Fatal(http.ListenAndServe(":8000", r))

}
