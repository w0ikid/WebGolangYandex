package main

import (
	"fmt"
	"net/http"
	_ "log"
)

const cookieName = "session_id"

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(cookieName)
	if err == http.ErrNoCookie {
		// Устанавливаем новый session_id
		newSessionID := "abc123" // Можно заменить на генератор UUID
		http.SetCookie(w, &http.Cookie{
			Name:  cookieName,
			Value: newSessionID,
			Path:  "/",
		})
		fmt.Fprintln(w, "Welcome!")
		// log.Printf("Welcome!")
	} else {
		// Если cookie уже есть
		fmt.Fprintln(w, "Welcome back!")
	}
}

func main() {
	http.HandleFunc("/", CookieHandler)
	// fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
