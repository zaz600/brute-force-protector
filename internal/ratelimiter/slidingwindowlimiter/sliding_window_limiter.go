package slidingwindowlimiter

import (
	"context"
	"sync"
	"time"

	"github.com/zaz600/brute-force-protector/internal/ratelimiter"
)

type SlidingWindowRateLimiter struct {
	ctx context.Context
	*sync.RWMutex
	db     map[string]*windowData
	window time.Duration
	limit  int64
}

func (r *SlidingWindowRateLimiter) Reset(key string) {
	r.Lock()
	defer r.Unlock()
	delete(r.db, key)
}

func (r *SlidingWindowRateLimiter) LimitReached(key string) bool {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[key]; !ok {
		r.db[key] = newWindowData(r.limit, r.window)
	}

	currentSize := r.db[key].currentSize()
	// log.Println(current, len(r.db[key].timestamps), r.db[key].timestamps)
	if currentSize < r.limit {
		r.db[key].add()
		return false
	}
	return true
}

func (r *SlidingWindowRateLimiter) cleanup() {
	r.Lock()
	defer r.Unlock()
	for k, v := range r.db {
		if v.currentSize() == 0 && time.Since(v.lastAccessTime) > 1*time.Minute {
			delete(r.db, k)
		}
	}
}

func NewSlidingWindowRateLimiter(ctx context.Context, window time.Duration, limit int64) ratelimiter.RateLimiter {
	limiter := &SlidingWindowRateLimiter{
		ctx:     ctx,
		RWMutex: &sync.RWMutex{},
		db:      make(map[string]*windowData),
		window:  window,
		limit:   limit,
	}

	go func() {
		for {
			select {
			case <-time.After(1 * time.Minute):
				limiter.cleanup()
			case <-ctx.Done():
				return
			}
		}
	}()
	return limiter
}
