package kafka

import (
	"context"
	"serviceB/cfg"
)

type manager struct {
	brokers string
	option  cfg.KafkaOption
}

type Manager interface {
	Write(ctx context.Context, message interface{}) (err error)
}

func NewManager(brokers string, op cfg.KafkaOption) Manager {
	return &manager{
		brokers: brokers,
		option:  op,
	}
}
