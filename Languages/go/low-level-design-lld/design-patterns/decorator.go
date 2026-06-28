package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Decorator Pattern: Logger wrapping Handler (middleware pattern).
// Each decorator wraps an http.Handler to add cross-cutting concerns
// without modifying the original handler.

// loggingMiddleware logs the request method, path, and duration.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[START] %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("[END]   %s %s (took %v)", r.Method, r.URL.Path, time.Since(start))
	})
}

// authMiddleware checks for an Authorization header.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Printf("[AUTH] Token present: %s...", token[:min(10, len(token))])
		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware catches panics and returns 500.
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] Recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// helloHandler is the base handler.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

// Chain applies middlewares in order (last applied = outermost).
func Chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Stack decorators: recovery -> logging -> auth -> handler
	handler := Chain(
		http.HandlerFunc(helloHandler),
		recoveryMiddleware,
		loggingMiddleware,
		authMiddleware,
	)

	fmt.Println("Server starting on :8080")
	fmt.Println("Test with: curl -H 'Authorization: Bearer token123' http://localhost:8080/world")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
