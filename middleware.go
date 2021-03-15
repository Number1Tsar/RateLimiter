package RateLimiter

import "net/http"

type RateLimitedMiddleware struct {
	Limiter RateLimiter
	Server  http.Handler
}

func (m RateLimitedMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !m.Limiter.AttemptRequest() {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Too many request attempts. Try after sometime."))
		return
	}
	m.Server.ServeHTTP(w, req)
}
