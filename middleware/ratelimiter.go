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

func (*FixedWindowRateLimiter) Limit() bool {
	// todo make sure timer is correct 
	mu.Lock()
	defer mu.Unlock()
	now := timer.Now()
	if !time || last - now > conf.duration {
		last = now
		passed = 0
	} 
	if passed >= conf.num {
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

func NewRateConfig(num int32, duration time.Duration) {
	res := new(RateConfig)
	res.num = num
	res.duration = duration
	return res
} 

func Example() {
	conf := NewRateConfig(60, timer.Duration(60) * timer.Second) 
	// one request per second
	limiter = NewFixedWindowRateLimiter(conf)
	
	_ = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			ratelimit.UnaryServerInterceptor(limiter),
		),
		grpc_middleware.WithStreamServerChain(
			ratelimit.StreamServerInterceptor(limiter),
		),
	) 
}
