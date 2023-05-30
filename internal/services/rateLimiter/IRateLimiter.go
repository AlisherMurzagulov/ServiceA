package rateLimiter

import "serviceB/redis"

type rateLimiter struct {
	redis redis.Manager
}

type IRateLimiter interface {
	AllowRequest(userId string) (allowRequest bool, err error)
}

func NewRateLimiter(redis redis.Manager) IRateLimiter {
	return &rateLimiter{
		redis: redis,
	}
}
