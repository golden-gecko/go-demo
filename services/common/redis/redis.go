package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	log "github.com/sirupsen/logrus"
)

var (
    Client *redis.Client
)

func Connect() error {
    log.Info("Connecting to Redis...")
	
	options := redis.Options {
        Addr:     "redis-1:6379",
        Password: "",
        DB:       0,
    }

    client := redis.NewClient(&options)

    log.Info("Connected to Redis")

	Client = client

    return nil
}

func Disconnect() {
    log.Info("Disconnecting from Redis...")

    Client.Close()

    log.Info("Disconnected from Redis")
}

func Set(key string, value string, expiration time.Duration) error {
	return Client.Set(context.Background(), key, value, expiration).Err()
}
