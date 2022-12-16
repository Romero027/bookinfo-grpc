package ratelimiter

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"google.golang.org/grpc"
	"time"
	"sync"
)



// rate is n / t
type RateConfig struct {
	num int32 // request num
	duration int64 // in ns
}

type FixedWindowRateLimiter struct {
	conf RateConfig
	last time.Time
	passed int32
	mu sync.Mutex
}

func (limiter *FixedWindowRateLimiter) Limit() bool {
	// todo make sure timer is correct 
	limiter.mu.Lock()
	defer limiter.mu.Unlock()
	now := time.Now()
	last := limiter.last
	dur := limiter.conf.duration
	num := limiter.conf.num
	if last.IsZero() || last.Sub(now).Nanoseconds() > dur {
		limiter.last = now
		limiter.passed = 0
	} 
	if limiter.passed >= num {
		return true
	} else {
		return false
	}
}

func NewFixedWindowRateLimiter(conf RateConfig) *FixedWindowRateLimiter {
	res := new(FixedWindowRateLimiter)
	res.conf = conf
	return res
}

func NewRateConfig(num int32, duration time.Duration) *RateConfig {
	res := new(RateConfig)
	res.num = num
	res.duration = duration.Nanoseconds()
	return res
} 

func Example() {
	conf := NewRateConfig(60, time.Duration(60) * time.Second) 
	// one request per second
	limiter := NewFixedWindowRateLimiter(*conf)
	
	_ = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			ratelimit.UnaryServerInterceptor(limiter),
		),
		grpc_middleware.WithStreamServerChain(
			ratelimit.StreamServerInterceptor(limiter),
		),
	) 
}
