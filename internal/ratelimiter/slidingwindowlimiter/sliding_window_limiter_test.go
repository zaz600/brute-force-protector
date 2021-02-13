package slidingwindowlimiter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/zaz600/brute-force-protector/internal/ratelimiter/slidingwindowlimiter"
)

func TestSlidingWindowRateLimiter_LimitReached(t *testing.T) {
	tests := []struct {
		name            string
		count           int
		expectedReached bool
	}{
		{
			name:            "not reached",
			count:           10,
			expectedReached: false,
		},
		{
			name:            "limit reached",
			count:           11,
			expectedReached: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, 10)
			for i := 0; i < tt.count; i++ {
				result := r.LimitReached("foo")
				if i == tt.count-1 {
					require.Equal(t, tt.expectedReached, result)
				} else {
					require.False(t, result)
				}
			}
		})
	}
}

func TestSlidingWindowRateLimiter_Reset(t *testing.T) {
	r := slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, 10)
	for i := 0; i < 11; i++ {
		r.LimitReached("foo")
	}
	result := r.LimitReached("foo")
	require.True(t, result)
	r.Reset("foo")
	result = r.LimitReached("foo")
	require.False(t, result)
}

func BenchmarkSlidingWindowRateLimiter_LimitReached(b *testing.B) {
	r := slidingwindowlimiter.NewSlidingWindowRateLimiter(time.Minute, 1000000)
	for i := 0; i < b.N; i++ {
		r.LimitReached("foo")
	}
}
