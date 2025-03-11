package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// Middleware для блокировки IP
func ipBlockerMiddleware(blockedIP string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)
		if clientIP == blockedIP {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Функция получения IP клиента с учётом прокси
func getClientIP(r *http.Request) string {
	// Проверяем X-Real-IP (обычно используется nginx)
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}
	// Проверяем X-Forwarded-For (может содержать несколько IP через запятую)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0]) // Берём первый IP (оригинальный)
	}
	// Получаем IP из r.RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // В редких случаях вернёт без изменений
	}
	return host
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Access granted")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", ipBlockerMiddleware("192.168.0.1", http.HandlerFunc(homeHandler)))

	// fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}
