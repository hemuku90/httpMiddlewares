package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello There!!!!!!!!!1")
}

//logginMiddleware
func logginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	fmt.Println("Middlewares")
	r := mux.NewRouter()
	r.Handle("/", logginMiddleware(http.HandlerFunc(indexHandler)))
	log.Fatal(http.ListenAndServe(":8000", r))
}
