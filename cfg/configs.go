package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Configs Конфигурация приложения
type Configs struct {
	Cache         Redis
	StopTimeoutMS int
	WebServer     WebServerConfig
	Kafka         KafkaSettings
	RateLimit     RateLimit
}

var AppConfigs = &Configs{}

type KafkaSettings struct {
	Brokers  string      `json:"Brokers"`
	Producer KafkaOption `json:"Producer"`
}

type KafkaOption struct {
	Topic          string `json:"Topic"`
	BatchBytes     int64  `json:"BatchBytes"`
	BatchTimeoutMS int    `json:"BatchTimeoutMS"`
	BatchSize      int    `json:"BatchSize"`
}

// Redis Настройка кэша приложения
type Redis struct {
	Host           string
	Port           int
	DatabaseNumber int
	Password       string
	Keys           struct {
		UserPayment RedisKey
		RateLimits  RedisKey
	}
}

// RedisKey models
type RedisKey struct {
	Key        string        `json:"key"`
	Expiration time.Duration `json:"expiration"`
}

// GetHostPort Возвращает <Host>:<Port>
func (r Redis) GetHostPort() string {
	return fmt.Sprintf("%v:%v", r.Host, r.Port)
}

// WebServerConfig Конфигурация веб сервера
type WebServerConfig struct {
	Port int
	GIN  GINConfig
}

type RateLimit struct {
	MaxRequestsPerSecond int
	RequestTimeWindow    int64
}

// GINConfig Конфигурация GIN
type GINConfig struct {
	ReleaseMode bool
	UseLogger   bool
	UseRecovery bool
}

// Setup Создает инициализированную конфигурацию
func Setup(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic("Config file reading error")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfigs)
	if err != nil {
		panic("Config file unmarshalling error")
	}
}
