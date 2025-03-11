package main

import (
	"fmt"
	"net/http"
	"log"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // передача упраления следующему обработчику
	})
}

func Middleware(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Middleware Test")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", loggingMiddleware(http.HandlerFunc(Middleware)))

	http.ListenAndServe(":8080", mux)
}