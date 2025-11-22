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

func ProcessVehicleCountByModel(client *redis.Client) error {
    models, err := roach.GetVehicleCountByModel()

    if err != nil {
		return err
    }

    for _, model := range models {
        err := client.Set(context.Background(), fmt.Sprintf("model_%s_%s", model.Brand, model.Model), strconv.Itoa(model.Count), time.Minute).Err()

        if err != nil {
			return err
        }
    }

	return nil
}

func ProcessVehicleCountByYear(client *redis.Client) error {
    years, err := roach.GetVehicleCountByYear()

    if err != nil {
		return err
    }

    for _, year := range years {
        err := client.Set(context.Background(), fmt.Sprintf("year_%d", year.Year), strconv.Itoa(year.Count), time.Minute).Err()

        if err != nil {
			return err
        }
    }

	return nil
}

func Sleep(interval int) {
	log.Info("Sleeping for ", interval, " ms")

	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func main() {
    if err := roach.Connect(); err != nil {
        log.Error(err)
        os.Exit(1)
    }

    defer roach.Disconnect()

	options := redis.Options{
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    }

    client := redis.NewClient(&options)

	defer client.Close()

    for {
        if err := ProcessVehicleCountByModel(client); err != nil {
			log.Error(err)
			os.Exit(1)
		}

        if err := ProcessVehicleCountByYear(client); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		Sleep(3000)
    }
}
