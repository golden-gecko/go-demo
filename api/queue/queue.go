package queue

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	kafka "github.com/segmentio/kafka-go"
)

var (
	Client  *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
	Context context.Context
	Cancel  context.CancelFunc

	KafkaClient *kafka.Conn
)

func InitRabbit() error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672")

	if err != nil {
		log.Println(err)
		return err
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		return err
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Println(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Client = conn
	Channel = ch
	Queue = q
	Context = ctx
	Cancel = cancel

	return nil
}

func InitKafka() error {
	topic := "my-topic"
	partition := 0

	connection, err := kafka.DialLeader(context.Background(), "tcp", "192.168.0.213:9092", topic, partition)

	if err != nil {
		log.Println(err)
		return err
	}

	connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

	KafkaClient = connection

	return nil
}

func DeinitKafka() error {
	err := KafkaClient.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeinitRabbit() error {
	Cancel()

	if err := Channel.Close(); err != nil {
		log.Println(err)
		return err
	}

	if err := Client.Close(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
