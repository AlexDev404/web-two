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

func main() {

	// Set the port to listen on
	port := 4000

	// the route pattern/endpoint/URL path
	http.Handle("/", loggingMiddleware(http.HandlerFunc(home)))

	// log the port we are listening on
	log.Print("starting server on :", port)

	// start a local web server
	err := http.ListenAndServe(fmt.Sprint(":", port), nil)
	log.Fatal(err)

}
