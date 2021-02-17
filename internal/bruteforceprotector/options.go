package bruteforceprotector

import (
	"context"

	"github.com/zaz600/brute-force-protector/internal/accesslist"
)

type ProtectorOption func(*BruteForceProtector)

// WithLoginLimit задает количество допустимых запросов для логина в минуту.
func WithLoginLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.loginLimit = maxCount
	}
}

// WithPasswordLimit задает количество допустимых запросов для пароля в минуту.
func WithPasswordLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.passwordLimit = maxCount
	}
}

// WithIPLimit задает количество допустимых запросов для IP в минуту.
func WithIPLimit(maxCount int64) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.ipLimit = maxCount
	}
}

// WithContext задает контекст.
func WithContext(ctx context.Context) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.ctx = ctx
	}
}

// WithBlackList устанавливает, какой использовать черный список.
func WithBlackList(l accesslist.AccessList) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.blackList = l
	}
}

// WithWhiteList устанавливает, какой использовать белый список.
func WithWhiteList(l accesslist.AccessList) ProtectorOption {
	return func(p *BruteForceProtector) {
		p.whiteList = l
	}
}
