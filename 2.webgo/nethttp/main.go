package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Data posted")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/submit", postHandler)
	http.ListenAndServe(":8080", nil)
}