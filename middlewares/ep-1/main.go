package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/hemuku90/middlewares/ep-1/pkg/services/api/handlers"
	middleware "github.com/hemuku90/middlewares/ep-1/pkg/services/middleware"
)

func main() {
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", handlers.HelloHandler)
	mux.HandleFunc("/v1/time", handlers.CurrentTimeHandler)
	//wrap entire mux with logger middleware
	wrappedMux := middleware.NewLogger(mux)
	log.Printf("server is listening at %s", addr)
	//use wrappedMux instead of mux as root handler
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
