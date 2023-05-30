package rateLimiter

import (
	"encoding/json"
	"fmt"
	"serviceB/cfg"
	"serviceB/internal/models"
	"time"
)

func (r *rateLimiter) AllowRequest(userId string) (allowRequest bool, err error) {
	redisKey := fmt.Sprintf("%s%s", cfg.AppConfigs.Cache.Keys.RateLimits, userId)
	if r.redis.Exits(redisKey) {
		var rateLimitModel models.RateLimiter

		keyData, keyTTL := r.redis.Get(redisKey)
		err = json.Unmarshal([]byte(keyData), &rateLimitModel)

		now := time.Now().Unix()
		allowRequest = now-rateLimitModel.FirstTimeRequest < cfg.AppConfigs.RateLimit.RequestTimeWindow && rateLimitModel.RequestCount < cfg.AppConfigs.RateLimit.MaxRequestsPerSecond

		rateLimitModel.RequestCount++
		r.redis.Delete(redisKey)
		r.redis.Set(redisKey, rateLimitModel, keyTTL)

		return
	}

	expiration := cfg.AppConfigs.Cache.Keys.RateLimits.Expiration
	newRateLimitModel := models.RateLimiter{
		RequestCount:     1,
		FirstTimeRequest: time.Now().Unix(),
	}

	r.redis.Set(redisKey, newRateLimitModel, expiration)

	return true, nil
}
