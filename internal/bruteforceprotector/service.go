package bruteforceprotector

import (
	"context"
	"time"

	"github.com/zaz600/brute-force-protector/internal/accesslist"
	"github.com/zaz600/brute-force-protector/internal/accesslist/memoryaccesslist"
	"github.com/zaz600/brute-force-protector/internal/ratelimiter"
	"github.com/zaz600/brute-force-protector/internal/ratelimiter/slidingwindowlimiter"
)

type BruteForceProtector struct {
	ctx context.Context

	blackList accesslist.AccessList
	whiteList accesslist.AccessList

	loginLimit    int64
	passwordLimit int64
	ipLimit       int64

	loginLimiter    ratelimiter.RateLimiter
	passwordLimiter ratelimiter.RateLimiter
	ipLimiter       ratelimiter.RateLimiter
}

func NewBruteForceProtector(opts ...ProtectorOption) *BruteForceProtector {
	// TODO 0 - no limit
	const (
		defaultMaxLoginAttempts    = 10
		defaultMaxPasswordAttempts = 100
		defaultMaxIPAttempts       = 1000
	)

	p := &BruteForceProtector{
		ctx:           context.Background(),
		loginLimit:    defaultMaxLoginAttempts,
		passwordLimit: defaultMaxPasswordAttempts,
		ipLimit:       defaultMaxIPAttempts,
	}

	for _, option := range opts {
		option(p)
	}

	p.loginLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(p.ctx, time.Minute, p.loginLimit)
	p.passwordLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(p.ctx, time.Minute, p.passwordLimit)
	p.ipLimiter = slidingwindowlimiter.NewSlidingWindowRateLimiter(p.ctx, time.Minute, p.ipLimit)

	if p.blackList == nil {
		p.blackList = memoryaccesslist.NewMemoryAccessList()
	}

	if p.whiteList == nil {
		p.whiteList = memoryaccesslist.NewMemoryAccessList()
	}

	return p
}

func (b *BruteForceProtector) Verify(ctx context.Context, login string, password string, ip string) bool {
	if inList := b.blackList.IsInList(ip); inList {
		return false
	}

	if inList := b.whiteList.IsInList(ip); inList {
		return true
	}

	return !b.limitReached(login, password, ip)
}

func (b *BruteForceProtector) limitReached(login string, password string, ip string) bool {
	loginCh := make(chan bool)
	passwordCh := make(chan bool)
	ipCh := make(chan bool)

	go func() {
		loginCh <- b.loginLimiter.LimitReached(login)
	}()

	go func() {
		passwordCh <- b.passwordLimiter.LimitReached(password)
	}()

	go func() {
		ipCh <- b.ipLimiter.LimitReached(ip)
	}()

	loginLimitReached := <-loginCh
	passwordLimitReached := <-passwordCh
	ipLimitReached := <-ipCh

	return loginLimitReached || passwordLimitReached || ipLimitReached
}

func (b *BruteForceProtector) ResetLogin(ctx context.Context, login string) {
	b.loginLimiter.Reset(login)
}

func (b *BruteForceProtector) ResetIP(ctx context.Context, ip string) {
	b.ipLimiter.Reset(ip)
}

func (b *BruteForceProtector) AddBlackList(ctx context.Context, networkCIDR string) error {
	return b.blackList.Add(networkCIDR)
}

func (b *BruteForceProtector) RemoveBlackList(ctx context.Context, networkCIDR string) error {
	return b.blackList.Remove(networkCIDR)
}

func (b *BruteForceProtector) BlackListItems(ctx context.Context) []string {
	return b.blackList.GetAll()
}

func (b *BruteForceProtector) AddWhiteList(ctx context.Context, networkCIDR string) error {
	return b.whiteList.Add(networkCIDR)
}

func (b *BruteForceProtector) RemoveWhiteList(ctx context.Context, networkCIDR string) error {
	return b.whiteList.Remove(networkCIDR)
}

func (b *BruteForceProtector) WhiteListItems(ctx context.Context) []string {
	return b.whiteList.GetAll()
}
