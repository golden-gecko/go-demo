package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/sirupsen/logrus"

	"consumer/queue"
)

type Coordinate struct {
	Latitude  float32 `bson:"Latitude"`
	Longitude float32 `bson:"Longitude"`
}

type Temperature struct {
	Location  string  `bson:"Location"`
	Value     float32 `bson:"Value"`
	Timestamp string  `bson:"Timestamp"`
}

type Transit struct {
	Plate     string     `bson:"Plate"`
	Location  Coordinate `bson:"Location"`
	Timestamp string     `bson:"Timestamp"`
}

func ProcessTemperature(writeAPI api.WriteAPIBlocking, body []byte) error {
	var temperature Temperature

	err := json.Unmarshal(body, &temperature)

	if err != nil {
		log.Error(err)
		return err
	}

	p := influxdb2.NewPointWithMeasurement("temperatue").
		AddTag("location", temperature.Location).
		AddTag("unit", "celsius").
		AddField("current", temperature.Value).
		SetTime(time.Now())

	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func ProcessTransit(body []byte) error {
	var transit Transit

	err := json.Unmarshal(body, &transit)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func main() {
	err := queue.Init()

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer queue.Deinit()

	client := influxdb2.NewClient("http://influx:8086", "my-super-secret-auth-token")

	if _, err := client.Health(context.Background()); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer client.Close()

	writeAPI := client.WriteAPIBlocking("my-org", "house")

	msgs1, err := queue.Channel.Consume(
		"temperatures",
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	msgs2, err := queue.Channel.Consume(
		"transits",
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	for {
		select {
		case d := <-msgs1:
			log.Debug("Received a message: %s", d.Body)

			if err := ProcessTemperature(writeAPI, d.Body); err != nil {
				if err := d.Nack(false, true); err != nil {
					log.Error(err)
					os.Exit(1)
				}
			} else {
				if err := d.Ack(false); err != nil {
					log.Error(err)
					os.Exit(1)
				}
			}
		case d := <-msgs2:
			log.Debug("Received a message: %s", d.Body)

			if err := ProcessTransit(d.Body); err != nil {
				if err := d.Nack(false, true); err != nil {
					log.Error(err)
					os.Exit(1)
				}
			} else {
				if err := d.Ack(false); err != nil {
					log.Error(err)
					os.Exit(1)
				}
			}
		}
	}
}
