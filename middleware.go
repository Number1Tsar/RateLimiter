package RateLimiter

import "net/http"

type RateLimitedMiddleware struct {
	Limiter RateLimiter
}

func (m *RateLimitedMiddleware) Serve(server http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !m.Limiter.Take() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many request attempts. Try after sometime."))
			return
		}
		server.ServeHTTP(w, req)
	})
}
