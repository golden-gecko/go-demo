package main

import (
	"log"

	"consumer/queue"
)

func main() {
	queue.Init()
	defer queue.Deinit()

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
		panic(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			d.Ack(false)
			// d.Nack(false, true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
