package queue

import (
	"context"
	"log"
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

func Init() error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")

	if err != nil {
		log.Println(err)
		return err
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		return err
	}

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	err = ch.Qos(10, 0, true)

	if err != nil {
		log.Println(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	Client = conn
	Channel = ch
	Queue = q
	Context = ctx
	Cancel = cancel

	return nil
}

func Deinit() error {
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
