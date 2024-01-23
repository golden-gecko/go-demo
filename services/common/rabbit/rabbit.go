package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var (
    Connection *amqp.Connection
    Channel    *amqp.Channel
    Queue       amqp.Queue
)

func Connect() error {
    connection, err := amqp.Dial("amqp://guest:guest@rabbit-1:5672")

    if err != nil {
        return err
    }

	Connection = connection

	log.Info("Connected to RabbitMQ")

    channel, err := connection.Channel()

    if err != nil {
        return err
    }

    Channel = channel

	if err = channel.Qos(10, 0, true); err != nil {
        return err
    }

    queues := []string{"items", "temperatures", "transits"}

    for _, queueName := range queues {
        if err := DeclareQueue(queueName); err != nil {
			return err
		}
    }

    return nil
}

func Disconnect() error {
    log.Info("Disconnecting from RabbitMQ...")

    if err := Channel.Close(); err != nil {
        return err
    }

    if err := Connection.Close(); err != nil {
        return err
    }

    log.Info("Disconnected from RabbitMQ")

    return nil
}

func DeclareQueue(name string) error {
	_, err := Channel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	return err
}

func ConsumeQueue(name string) (<-chan amqp.Delivery, error) {
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
