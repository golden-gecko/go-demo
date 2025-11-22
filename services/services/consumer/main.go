package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rabbitmq/amqp091-go"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/sirupsen/logrus"

	"services/common/model"
	"services/common/queue"
)

func ProcessTemperature(writeAPI api.WriteAPIBlocking, body []byte) error {
    var temperature model.Temperature

    err := json.Unmarshal(body, &temperature)

    if err != nil {
        return err
    }

    point := influxdb2.NewPointWithMeasurement("temperatue").
		AddTag("location", temperature.Location).
        AddTag("unit", "celsius").
        AddField("current", temperature.Value).
        SetTime(time.Now()) // TODO: Move timestamp to producer.

    if err := writeAPI.WritePoint(context.Background(), point); err != nil {
        return err
    }

    return nil
}

func ProcessTransit(body []byte) error {
    var transit model.Transit

    err := json.Unmarshal(body, &transit)

    if err != nil {
        return err
    }

    return nil
}

func CreateQueue(name string) (<-chan amqp091.Delivery, error) {
	return queue.Channel.Consume(
		name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func ProcessTemperatureQueue(writeAPI api.WriteAPIBlocking, msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		// log.Info("Received a message: %s", d.Body)

		if err := ProcessTemperature(writeAPI, d.Body); err != nil {
			log.Error(err)

			if err := d.Nack(false, true); err != nil {
				panic(err)
			}
		} else if err := d.Ack(false); err != nil {
			panic(err)
		}
	}

	writeAPI.Flush(context.Background())
}

func ProcessTransitQueue(writeAPI api.WriteAPIBlocking, msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		// log.Info("Received a message: %s", d.Body)

		if err := ProcessTransit(d.Body); err != nil {
			log.Error(err)

			if err := d.Nack(false, true); err != nil {
				panic(err)
			}
		} else if err := d.Ack(false); err != nil {
			panic(err)
		}
	}
}

func main() {
    if err := queue.Init(); err != nil {
		panic(err)
    }

    defer queue.Deinit()

	log.Info("Connected to RabbitMQ")

    client := influxdb2.NewClientWithOptions(
		"http://influx:8086",
		"my-super-secret-auth-token",
		influxdb2.DefaultOptions().SetBatchSize(10))

    if _, err := client.Health(context.Background()); err != nil {
		panic(err)
    }

    defer client.Close()

	log.Info("Connected to InfluxDB")

    writeAPI := client.WriteAPIBlocking("my-org", "house")

	msgs1, err := CreateQueue("temperatures")

    if err != nil {
		panic(err)
    }

	log.Info("Queue temperatures created")

	msgs2, err := CreateQueue("transits")

    if err != nil {
		panic(err)
    }

	log.Info("Queue transits created")
	log.Info("Waiting for messages...")

	var forever chan struct {}

	go ProcessTemperatureQueue(writeAPI, msgs1)
	go ProcessTransitQueue(writeAPI, msgs2)

	<-forever
}
