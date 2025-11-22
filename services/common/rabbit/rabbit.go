package rabbit

import (
	"context"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

var (
    connection *amqp.Connection
    channel    *amqp.Channel
)

func Connect() error {
    log.Info("Connecting to RabbitMQ...")

	user := os.Getenv("RABBIT_USER")
	password := os.Getenv("RABBIT_PASSWORD")
	host := os.Getenv("RABBIT_HOST")
	port := os.Getenv("RABBIT_PORT")

    connection_, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port))

    if err != nil {
        return err
    }

	connection = connection_

	log.Info("Connected to RabbitMQ")

    channel_, err := connection.Channel()

    if err != nil {
        return err
    }

    channel = channel_

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

    if err := channel.Close(); err != nil {
        return err
    }

    if err := connection.Close(); err != nil {
        return err
    }

    log.Info("Disconnected from RabbitMQ")

    return nil
}

func DeclareQueue(name string) error {
	_, err := channel.QueueDeclare(
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
	return channel.Consume(
		name,  // name
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func Write(queueName string, data []byte) error {
	publishing := amqp.Publishing {
		ContentType:  "text/plain",
		Body:         data,
		DeliveryMode: 2,
	}

    return channel.PublishWithContext(
        context.Background(),
        "",
        queueName,
        false,      // mandatory
        false,      // immediate
		publishing,
    )
}
