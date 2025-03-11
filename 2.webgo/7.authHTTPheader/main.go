package main

import (
	"fmt"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		
		if authHeader != "Bearer valid_token" {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authorized access")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", authMiddleware(http.HandlerFunc(protectedHandler)))
	// fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}