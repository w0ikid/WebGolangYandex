package main

import (
	"fmt"
	"net/http"
	"io"
)

func startServer(address string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})

	http.ListenAndServe(address, nil)
}

func sendRequest(url string) (string, error) {
	resp, err := http.Get("http://" + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// go startServer()

	resp, err := sendRequest("http://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resp)
}