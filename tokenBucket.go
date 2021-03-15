package RateLimiter

import (
	"sync"
	"time"
)

type bucket struct {
	sync.Mutex
	capacity       int
	refillQuantum  int
	refillInterval time.Duration
	currentToken   int
}

// NewBucket initializes TokenBucket based rate Limiter with capacity, refill amount and refill interval.
func NewBucket(capacity int, refillQuantum int, refillInterval time.Duration) RateLimiter {
	b := &bucket{
		capacity:       capacity,
		refillQuantum:  refillQuantum,
		refillInterval: refillInterval,
		currentToken:   capacity,
	}
	periodicallyRefill(b)
	return b
}

// Take checks if enough tokens are available to service request. If the tokens are present it signals true otherwise
// returns false
func (b *bucket) Take() bool {
	b.Lock()
	defer b.Unlock()
	if b.currentToken <= 0 {
		return false
	}
	b.currentToken -= 1
	return true
}

func periodicallyRefill(b *bucket) {
	tick := time.NewTicker(b.refillInterval)
	go func() {
		for {
			select {
			case <-tick.C:
				b.Lock()
				if b.currentToken+b.refillQuantum >= b.capacity {
					b.currentToken = b.capacity
				} else {
					b.currentToken += b.refillQuantum
				}
				b.Unlock()
			default:
			}
		}
	}()
}
