package bruteforceprotector

import (
	"time"

	"github.com/zaz600/brute-force-protector/internal/accesslist/redisaccesslist"
	"github.com/zaz600/brute-force-protector/internal/ratelimiter/slidingwindowlimiter"
)

type ProtectorOption func(*BruteForceProtector)

func WithLoginLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.loginLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, maxCount)
	}
}

func WithPasswordLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.passwordLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, maxCount)
	}
}

func WithIPLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.ipLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, maxCount)
	}
}

func WithRedis(host string) ProtectorOption {
	return func(p *BruteForceProtector) {
		if host != "" {
			p.blackList = redisaccesslist.NewRedisAccessList("blacklist", host)
			p.whiteList = redisaccesslist.NewRedisAccessList("whitelist", host)
		}
	}
}
