package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Number1Tsar/RateLimiter"
	"github.com/Number1Tsar/RateLimiter/internal/tokenbucket"
	"github.com/gorilla/mux"
)

func dummyService(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully handled request."))
}

func main() {
	rateLimiter := tokenbucket.NewBucket(100, 10, time.Second)
	m := mux.NewRouter()
	service := RateLimiter.RateLimitedMiddleware{
		Limiter: rateLimiter,
		Server:  http.HandlerFunc(dummyService),
	}
	m.Handle("/a", service)
	m.Handle("/b", http.HandlerFunc(dummyService))
	log.Fatal(http.ListenAndServe(":8000", m))
}
