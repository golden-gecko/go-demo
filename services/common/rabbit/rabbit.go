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
    connection, err := amqp.Dial("amqp://guest:guest@rabbit-1:5672")

    if err != nil {
        return err
    }

	log.Info("Connected to RabbitMQ")

    ch, err := connection.Channel()

    if err != nil {
        return err
    }

	if err = ch.Qos(10, 0, true); err != nil {
        return err
    }

    queues := []string{"items", "temperatures", "transits"}

    for _, queueName := range queues {
        if err := DeclareQueue(queueName); err != nil {
			return err
		}
    }

    Client = connection
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
