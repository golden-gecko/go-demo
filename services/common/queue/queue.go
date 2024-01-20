package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
    Client  *amqp.Connection
    Channel *amqp.Channel
    Queue   amqp.Queue
    Context context.Context
    Cancel  context.CancelFunc
)

func Connect() error {
    conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")

    if err != nil {
        return err
    }

    ch, err := conn.Channel()

    if err != nil {
        return err
    }

    err = ch.Qos(10, 0, true)

    queues := []string{"items", "temperatures", "transits"}

    for _, queueName := range queues {
        _, err := ch.QueueDeclare(
            queueName, // name
            true,      // durable
            false,     // delete when unused
            false,     // exclusive
            false,     // no-wait
            nil,       // arguments
        )

        if err != nil {
            return err
        }
    }

    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

    Client = conn
    Channel = ch
    Context = ctx
    Cancel = cancel

    return nil
}

func Disconnect() error {
    Cancel()

    if err := Channel.Close(); err != nil {
        return err
    }

    if err := Client.Close(); err != nil {
        return err
    }

    return nil
}

func CreateQueue(name string) (<-chan amqp.Delivery, error) {
	return Channel.Consume(
		name,  // name
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}
