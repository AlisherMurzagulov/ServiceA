package main

import (
	"fmt"
	"os"
	"os/signal"
	"serviceB/cfg"
	"serviceB/internal/http"
	"serviceB/kafka"
	"serviceB/redis"
	"syscall"
)

func init() {
	cfg.Setup("config.json")
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	redis := redis.NewManager()
	kafka := kafka.NewManager(cfg.AppConfigs.Kafka.Brokers, cfg.AppConfigs.Kafka.Producer)

	server, err := http.NewServer(redis, kafka)
	if err != nil {
		panic(fmt.Sprintf("create server error: %v", err))
	}

	startServerError := server.Start()
	var stopReason string
	select {
	case err = <-startServerError:
		stopReason = fmt.Sprintf("start server error: %v", err)
	case qs := <-quit:
		stopReason = fmt.Sprintf("received signal %s", qs.String())
	}

	fmt.Printf("%s\nshutting down server...\n", stopReason)
	err = server.Stop()
	if err != nil {
		fmt.Printf("stop server error: %v\n", err)
		return
	}
	fmt.Println("server stopped")
}
