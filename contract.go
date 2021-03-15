package RateLimiter

// RateLimiter is used to control the rate at which server consumes any request.
type RateLimiter interface {
	Take() bool
}
