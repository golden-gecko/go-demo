package main

import (
	"fmt"
	"strconv"
	"time"

	"services/common"
	"services/common/redis"
	"services/common/roach"
)

func ProcessVehicleCountByModel(interval int) {
	if err := roach.Connect(); err != nil {
		panic(err)
    }

    defer roach.Disconnect()

    models, err := roach.GetVehicleCountByModel()

    if err != nil {
		panic(err)
    }

	for {
		for _, model := range models {
			err := redis.Set(fmt.Sprintf("model_%s_%s", model.Brand, model.Model), strconv.Itoa(model.Count), time.Minute)

			if err != nil {
				panic(err)
			}
		}

		common.Sleep(interval)
	}
}

func ProcessVehicleCountByYear(interval int) {
	if err := roach.Connect(); err != nil {
		panic(err)
    }

    defer roach.Disconnect()

    years, err := roach.GetVehicleCountByYear()

    if err != nil {
		panic(err)
    }

	for {
		for _, year := range years {
			err := redis.Set(fmt.Sprintf("year_%d", year.Year), strconv.Itoa(year.Count), time.Minute)

			if err != nil {
				panic(err)
			}
		}

		common.Sleep(interval)
	}
}

func main() {
    if err := redis.Connect(); err != nil {
		panic(err)
	}

	defer redis.Disconnect()

	var forever chan struct {}

	go ProcessVehicleCountByModel(3000)
	go ProcessVehicleCountByYear(6000)

	<- forever
}
