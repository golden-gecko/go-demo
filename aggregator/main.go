package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	log "github.com/sirupsen/logrus"

	"aggregator/roach"
)

func ProcessVehicleCountByModel(client *redis.Client) {
    models, err := roach.GetVehicleCountByModel()

    if err != nil {
        log.Error(err)
        os.Exit(1)
    }

    for _, model := range models {
        err := client.Set(context.Background(), fmt.Sprintf("%s %s", model.Brand, model.Model), strconv.Itoa(model.Count), 0).Err()

        if err != nil {
        	log.Error(err)
			os.Exit(1)
        }
    }
}

func ProcessVehicleCountByYear(client *redis.Client) {
    years, err := roach.GetVehicleCountByYear()

    if err != nil {
        log.Error(err)
        os.Exit(1)
    }

    for _, year := range years {
        err := client.Set(context.Background(), strconv.Itoa(year.Year), strconv.Itoa(year.Count), 0).Err()

        if err != nil {
        	log.Error(err)
			os.Exit(1)
        }
    }
}

func main() {
    if err := roach.Connect(); err != nil {
        log.Error(err)
        os.Exit(1)
    }

    defer roach.Disconnect()

    client := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    interval := 3000

    for {
        ProcessVehicleCountByModel(client)
        ProcessVehicleCountByYear(client)

        log.Debug("Sleeping for", interval, "ms")

        time.Sleep(time.Duration(interval) * time.Millisecond)
    }
}
