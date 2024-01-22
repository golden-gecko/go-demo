package main

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"

	log "github.com/sirupsen/logrus"

	"services/common/influx"
	"services/common/model"
	"services/common/mongo"
	"services/common/rabbit"
)

func ProcessItem(body []byte) error {
    var item model.Item

    err := json.Unmarshal(body, &item)

    if err != nil {
        return err
    }

    return nil
}

func ProcessTemperature(body []byte) error {
    var temperature model.Temperature

    err := json.Unmarshal(body, &temperature)

    if err != nil {
        return err
    }

    if err := influx.WriteTemperature(temperature.Location, temperature.Value, temperature.Timestamp); err != nil {
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

func ProcessItemQueue(msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		// log.Info("Received a message: %s", d.Body)

		if err := ProcessItem(d.Body); err != nil {
			log.Error(err)

			if err := d.Nack(false, true); err != nil {
				panic(err)
			}
		} else if err := d.Ack(false); err != nil {
			panic(err)
		}
	}
}

func ProcessTemperatureQueue(msgs <-chan amqp091.Delivery) {
	for d := range msgs {
		// log.Info("Received a message: %s", d.Body)

		if err := ProcessTemperature(d.Body); err != nil {
			log.Error(err)

			if err := d.Nack(false, true); err != nil {
				panic(err)
			}
		} else if err := d.Ack(false); err != nil {
			panic(err)
		}
	}
}

func ProcessTransitQueue(msgs <-chan amqp091.Delivery) {
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
	if err := influx.Connect(); err != nil {
		panic(err);
	}

	defer influx.Disconnect();

	//

	if err := mongo.Connect(); err != nil {
		panic(err);
	}

	defer mongo.Disconnect();

	//

    if err := rabbit.Connect(); err != nil {
		panic(err)
    }

    defer rabbit.Disconnect()

	//

	msgs1, err := rabbit.CreateQueue("items")

    if err != nil {
		panic(err)
    }

	log.Info("Queue items created")

	msgs2, err := rabbit.CreateQueue("temperatures")

    if err != nil {
		panic(err)
    }

	log.Info("Queue temperatures created")

	msgs3, err := rabbit.CreateQueue("transits")

    if err != nil {
		panic(err)
    }

	log.Info("Queue transits created")

	//

	log.Info("Waiting for messages...")

	var forever chan struct {}

	go ProcessItemQueue(msgs1)
	go ProcessTemperatureQueue(msgs2)
	go ProcessTransitQueue(msgs3)

	<-forever
}
