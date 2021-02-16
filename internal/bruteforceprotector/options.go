package bruteforceprotector

import (
	"context"

	"github.com/zaz600/brute-force-protector/internal/accesslist"
)

type ProtectorOption func(*BruteForceProtector)

func WithLoginLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.loginLimit = maxCount
	}
}

func WithPasswordLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.passwordLimit = maxCount
	}
}

func WithIPLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.ipLimit = maxCount
	}
}

func WithContext(ctx context.Context) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.ctx = ctx
	}
}

func WithBlackList(l accesslist.AccessList) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.blackList = l
	}
}

func WithWhiteList(l accesslist.AccessList) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.whiteList = l
	}
}
