package main

import (
	"fmt"
	"strconv"
	"time"

	"services/common"
	"services/common/metrics"
	"services/common/redis"
	"services/common/roach"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func ProcessVehicleCountByModel(recordsProcessed prometheus.Counter, interval int) {
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

			recordsProcessed.Inc()
		}

		common.Sleep(interval)
	}
}

func ProcessVehicleCountByYear(recordsProcessed prometheus.Counter, interval int) {
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

			recordsProcessed.Inc()
		}

		common.Sleep(interval)
	}
}

func main() {
    if err := redis.Connect(); err != nil {
		panic(err)
	}

	defer redis.Disconnect()

	recordsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "aggregator_records_processed",
	})

	var forever chan struct {}

	go metrics.Serve(9105)

	go ProcessVehicleCountByModel(recordsProcessed, 3000)
	go ProcessVehicleCountByYear(recordsProcessed, 6000)

	<- forever
}
