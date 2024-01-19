package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"services/common"
	"services/common/sql"
)

func ProcessVehicleCountByModel(client *redis.Client, interval int) {
	if err := sql.Connect(); err != nil {
		panic(err)
    }

    defer sql.Disconnect()

    models, err := sql.GetVehicleCountByModel()

    if err != nil {
		panic(err)
    }

	for {
		for _, model := range models {
			err := client.Set(context.Background(), fmt.Sprintf("model_%s_%s", model.Brand, model.Model), strconv.Itoa(model.Count), time.Minute).Err()

			if err != nil {
				panic(err)
			}
		}

		common.Sleep(interval)
	}
}

func ProcessVehicleCountByYear(client *redis.Client, interval int) {
	if err := sql.Connect(); err != nil {
		panic(err)
    }

    defer sql.Disconnect()

    years, err := sql.GetVehicleCountByYear()

    if err != nil {
		panic(err)
    }

	for {
		for _, year := range years {
			err := client.Set(context.Background(), fmt.Sprintf("year_%d", year.Year), strconv.Itoa(year.Count), time.Minute).Err()

			if err != nil {
				panic(err)
			}
		}

		common.Sleep(interval)
	}
}

func main() {
	options := redis.Options {
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    }

    client := redis.NewClient(&options)

	defer client.Close()

	var forever chan struct {}

	go ProcessVehicleCountByModel(client, 30000)
	go ProcessVehicleCountByYear(client, 60000)

	<- forever
}
