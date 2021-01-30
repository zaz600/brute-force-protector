package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter interface {
	LimitReached(key string) bool
	Reset(key string)
}

type SlidingWindowRateLimiter struct {
	*sync.RWMutex
	db     map[string][]int64
	window time.Duration
	limit  int64
}

func (r *SlidingWindowRateLimiter) Reset(key string) {
	r.Lock()
	defer r.Unlock()
	r.db[key] = make([]int64, 0, r.limit)
}

func (r *SlidingWindowRateLimiter) LimitReached(key string) bool {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[key]; !ok {
		r.db[key] = make([]int64, 0, r.limit)
	}
	current := r.getCountInWindow(key)
	// log.Println(current, len(r.db[key]), r.db[key])
	if current < r.limit {
		r.db[key] = append(r.db[key], time.Now().UnixNano())
		return false
	}
	return true
}

func (r *SlidingWindowRateLimiter) getCountInWindow(key string) int64 {
	var count int64
	firstLeftEl := -1
	windowLeft := time.Now().UnixNano() - r.window.Nanoseconds()
	for i, value := range r.db[key] {
		if value >= windowLeft {
			count++
			if firstLeftEl == -1 {
				firstLeftEl = i
			}
		}
	}
	// TODO унести в отдельную
	if firstLeftEl > 0 {
		r.db[key] = r.db[key][firstLeftEl : len(r.db[key])-1]
	}
	return count
}

func NewSlidingWindowRateLimiter(window time.Duration, limit int64) RateLimiter {
	return &SlidingWindowRateLimiter{
		RWMutex: &sync.RWMutex{},
		db:      make(map[string][]int64),
		window:  window,
		limit:   limit,
	}
}
