package main

import (
	"net/http"
	"testing"
	"time"
)

func TestCookieHandler(t *testing.T) {
	go main()
	time.Sleep(time.Second)

	client := &http.Client{}

	// Test without cookie
	resp, err := client.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	// Test with cookie
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "test"})

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("server does not respond")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}