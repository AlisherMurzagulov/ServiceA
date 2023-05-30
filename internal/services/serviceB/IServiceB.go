package serviceB

import (
	"github.com/gin-gonic/gin"
	"serviceB/internal/models"
	"serviceB/internal/services/rateLimiter"
	"serviceB/kafka"
	"serviceB/redis"
)

type seviceB struct {
	redis   redis.Manager
	kafka   kafka.Manager
	limiter rateLimiter.IRateLimiter
}

type IServiceB interface {
	Transfer(ctx *gin.Context, payment models.Request, userId string) (err error)
}

func NewServiceB(redis redis.Manager, kafka kafka.Manager, limiter rateLimiter.IRateLimiter) IServiceB {
	return &seviceB{
		redis:   redis,
		kafka:   kafka,
		limiter: limiter,
	}
}
