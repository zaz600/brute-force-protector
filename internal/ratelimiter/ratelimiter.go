package ratelimiter

type RateLimiter interface {
	LimitReached(key string) bool
	Reset(key string)
}
