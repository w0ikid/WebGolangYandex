package main

import (
	"fmt"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.ListenAndServe(":8080", nil)
}