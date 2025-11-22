package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"

	log "github.com/sirupsen/logrus"
)

var (
    client *redis.Client
)

func Connect() error {
    log.Info("Connecting to Redis...")
	
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	options := redis.Options {
        Addr:     fmt.Sprintf("%s:%s", host, port),
        Password: "",
        DB:       0,
    }

    client = redis.NewClient(&options)

    log.Info("Connected to Redis")

    return nil
}

func Disconnect() {
    log.Info("Disconnecting from Redis...")

    client.Close()

    log.Info("Disconnected from Redis")
}

func Set(key string, value string, expiration time.Duration) error {
	return client.Set(context.Background(), key, value, expiration).Err()
}
