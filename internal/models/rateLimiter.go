package models

type RateLimiter struct {
	FirstTimeRequest int64
	RequestCount     int
}
