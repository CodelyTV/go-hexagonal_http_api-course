package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	mux := http.NewServeMux()

	healthHandler := http.HandlerFunc(healthHandler)
	mux.Handle("/health", recoveryMiddleware(healthHandler))

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("everything is ok!"))
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Middleware] %s panic recovered:\n%s\n",
					time.Now().Format("2006/01/02 - 15:04:05"), err)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
