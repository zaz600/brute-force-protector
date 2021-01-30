package bruteforceprotector

import (
	"time"

	"github.com/zaz600/brute-force-protector/internal/accesslist"
	"github.com/zaz600/brute-force-protector/internal/ratelimiter"
)

type BruteForceProtector struct {
	/*
	   	не более loginLimit = 10 попыток в минуту для данного логина.
	       не более M = 100 попыток в минуту для данного пароля (защита от обратного brute-force).
	       не более K = 1000 попыток в минуту для данного IP (число большое, т.к. NAT).
	*/
	// TODO ctx!

	blackList accesslist.AccessList
	whiteList accesslist.AccessList

	loginLimiter    ratelimiter.RateLimiter
	passwordLimiter ratelimiter.RateLimiter
	ipLimiter       ratelimiter.RateLimiter
}

func NewBruteForceProtector(n, m, k int64) *BruteForceProtector {
	return &BruteForceProtector{
		blackList: accesslist.NewMemoryAccessList(),
		whiteList: accesslist.NewMemoryAccessList(),

		loginLimiter:    ratelimiter.NewSlidingWindowRateLimiter(time.Minute, n),
		passwordLimiter: ratelimiter.NewSlidingWindowRateLimiter(time.Minute, m),
		ipLimiter:       ratelimiter.NewSlidingWindowRateLimiter(time.Minute, k),
	}
}

func (b *BruteForceProtector) Verify(login string, password string, ip string) bool {
	// TODO что делать если ip в обоих листах?

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

func (b *BruteForceProtector) ResetLogin(login string) {
	b.loginLimiter.Reset(login)
}

func (b *BruteForceProtector) ResetIP(ip string) {
	b.ipLimiter.Reset(ip)
}

func (b *BruteForceProtector) AddBlackList(networkCIDR string) error {
	return b.blackList.Add(networkCIDR)
}

func (b *BruteForceProtector) RemoveBlackList(networkCIDR string) {
	b.blackList.Remove(networkCIDR)
}

func (b *BruteForceProtector) AddWhiteList(networkCIDR string) error {
	return b.whiteList.Add(networkCIDR)
}

func (b *BruteForceProtector) RemoveWhiteList(networkCIDR string) {
	b.whiteList.Remove(networkCIDR)
}
