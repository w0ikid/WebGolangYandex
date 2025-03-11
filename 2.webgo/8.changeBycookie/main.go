package main

import (
	"fmt"
	"net/http"
)

func languageHandler(w http.ResponseWriter, r *http.Request) {
	const defaultLang = "en"
	supportedLanguages := map[string]string{
		"en": "Hello!",
		"ru": "Привет!",
	}

	// Получаем cookie lang
	cookie, err := r.Cookie("lang")
	if err != nil || cookie.Value == "" {
		// Если cookie нет, устанавливаем английский по умолчанию
		http.SetCookie(w, &http.Cookie{
			Name:  "lang",
			Value: defaultLang,
			Path:  "/",
		})
		fmt.Fprintln(w, supportedLanguages[defaultLang])
		return
	}

	// Проверяем, поддерживается ли язык
	if msg, exists := supportedLanguages[cookie.Value]; exists {
		fmt.Fprintln(w, msg)
	} else {
		fmt.Fprintln(w, supportedLanguages[defaultLang])
	}
}

func main() {
	http.HandleFunc("/", languageHandler)
	http.ListenAndServe(":8080", nil)
}