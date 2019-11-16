package routes

import "net/http"

// Home is a handler for the Home route
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Welcome to the Home page"))
}

// About is a handler for the About route
func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Welcome to the About page"))
}

// NotFound is a handler for the NotFound route
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("404 - Page not found"))
}
