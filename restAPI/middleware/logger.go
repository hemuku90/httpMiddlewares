package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

//Middlerware Type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Logging by middleware
			start := time.Now()
			defer func() {
				log.Debugf("Incomming request %s %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, r.URL.Path, time.Since(start))
			}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

//Chain multiple middlewares
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//Log Logmiddleware
// func Logmiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		loggger.Debugf("Incomming request %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
// 		next.ServeHTTP(w, r)
// 	})
// }

// //Loggin middleware
// func Logging(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path)
// 		log.Println(r
// 		f(w, r)
// 	}
// }
