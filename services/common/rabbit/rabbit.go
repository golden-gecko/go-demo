package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var (
    Client  *amqp.Connection
    Channel *amqp.Channel
    Queue   amqp.Queue
)

func Connect() error {
    conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")

    if err != nil {
        return err
    }

	log.Info("Connected to RabbitMQ")

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

    Client = conn
    Channel = ch

    return nil
}

func Disconnect() error {
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
