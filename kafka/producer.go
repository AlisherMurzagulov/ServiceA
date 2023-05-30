package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func (m *manager) Write(ctx context.Context, message interface{}) (err error) {
	w := &kafka.Writer{
		Addr:         kafka.TCP(m.brokers),
		Topic:        m.option.Topic,
		Balancer:     &kafka.RoundRobin{},
		BatchSize:    m.option.BatchSize,
		BatchBytes:   m.option.BatchBytes,
		BatchTimeout: time.Duration(m.option.BatchTimeoutMS) * time.Millisecond,
		Compression:  kafka.Gzip,
	}

	//writerCollector := collector.NewWriterCollector(w)
	//prometheus.MustRegister(writerCollector)

	messageBytes, err := json.Marshal(&message)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = w.WriteMessages(ctx,
		kafka.Message{
			Value: messageBytes,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer func() {
		fmt.Println("kafka writer close")
		fmt.Println("stats:")
		stats, _ := json.Marshal(w.Stats())
		fmt.Println(stats)
		err = w.Close()
	}()

	return nil
}
