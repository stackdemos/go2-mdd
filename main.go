package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"golang-backend/consts"
	"golang-backend/controllers"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI, r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func withUptime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, consts.UptimeKey, uptime())

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", controllers.Status)
	r.HandleFunc("/status", controllers.Status)

	r.Use(withLogging)
	r.Use(withUptime)

	// Bind to a port and pass our router in
	log.Println("Listening 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
