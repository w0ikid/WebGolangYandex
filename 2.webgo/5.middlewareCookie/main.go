package main

import (
  "fmt"
  "log"
  "net/http"
)

func authMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    _, err := r.Cookie("user_id")
    if err != nil {
      http.Redirect(w, r, "/login", http.StatusFound)
            return
    }
    next.ServeHTTP(w, r)
  })
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  cookie := &http.Cookie{
    Name:  "user_id",
    Value: "123",
    Path:  "/",
  }
  http.SetCookie(w, cookie)
  fmt.Fprintln(w, "Please log in")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Access granted")
}

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/login", loginHandler)

  mux.Handle("/", authMiddleware(http.HandlerFunc(homeHandler)))

//   log.Println("Server running on :8080")
  if err := http.ListenAndServe(":8080", mux); err != nil {
    log.Fatal("Server error ", err)
  }
}
