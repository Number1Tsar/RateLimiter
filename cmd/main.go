package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Number1Tsar/RateLimiter"
)

func dummyService(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully handled request."))
}

func main() {
	rateLimiter := RateLimiter.NewBucket(100, 10, time.Second)
	service := RateLimiter.RateLimitedMiddleware{
		Limiter: rateLimiter,
	}
	http.Handle("/", service.Serve(http.HandlerFunc(dummyService)))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
