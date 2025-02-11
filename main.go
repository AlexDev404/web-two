// Filename: main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// A handler function named 'home'
func home(rz http.ResponseWriter, r *http.Request) {
	rz.Write([]byte("Hello from Flash"))
	// fmt.Print("Request from User Agent: ", r.UserAgent())
}

// A basic middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Print("Method: ", r.Method, " URL: ", r.URL.Path, " Time: ", time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// Admin middleware
func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is an admin
		if r.Header.Get("X-Admin") != "true" {
			// If they are not, return a 403 Forbidden
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// If they are, call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Set the port to listen on
	port := 4000

	// the route pattern/endpoint/URL path
	mux.HandleFunc("/", home)
	mux.Handle("/admin", adminMiddleware(http.HandlerFunc(home))) // adminMiddleware is a middleware

	// log the port we are listening on
	log.Print("starting server on :", port)

	// start a local web server
	err := http.ListenAndServe(fmt.Sprint(":", port), loggingMiddleware(mux))
	log.Fatal(err)

}
