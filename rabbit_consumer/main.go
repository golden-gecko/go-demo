package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"consumer/queue"
)

func main() {
	err := queue.Init()

	if err != nil {
		os.Exit(1)
	}

	defer queue.Deinit()

	//
	client := influxdb2.NewClient("http://influx:8086", "my-super-secret-auth-token")

	if _, err := client.Health(context.Background()); err != nil {
		os.Exit(1)
	}

	defer client.Close()

	writeAPI := client.WriteAPIBlocking("my-org", "my-bucket")
	//

	msgs, err := queue.Channel.Consume(
		queue.Queue.Name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		//
		p := influxdb2.NewPointWithMeasurement("stat").
			AddTag("unit", "temperature").
			AddField("avg", rand.Float32()*10).
			AddField("max", rand.Float32()*20).
			SetTime(time.Now())

		if err := writeAPI.WritePoint(context.Background(), p); err != nil {
			log.Panicln(err)
		}
		//

		if err := d.Ack(false); err != nil {
			os.Exit(1)
		}

		// d.Nack(false, true)
	}
}
