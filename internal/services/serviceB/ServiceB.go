package serviceB

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"serviceB/cfg"
	"serviceB/internal/models"
)

func (s *seviceB) Transfer(ctx *gin.Context, payment models.Request, userId string) (err error) {

	if ok, _ := s.limiter.AllowRequest(userId); !ok {
		return fmt.Errorf("rate limit exceeded: %v", userId)
	}

	userKey := fmt.Sprintf("%s%s:%s", cfg.AppConfigs.Cache.Keys.UserPayment, userId, payment.ID)
	expiration := cfg.AppConfigs.Cache.Keys.UserPayment.Expiration
	if s.redis.Exits(userKey) {
		return fmt.Errorf("user has exiting payment : %v", payment.ID)
	}

	err = s.kafka.Write(ctx, payment)
	if err != nil {
		return fmt.Errorf("failed to send request to Kafka: %v", err)
	}

	s.redis.Set(userKey, payment, expiration)

	return
}
