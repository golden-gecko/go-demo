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

func Init() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.10.23:5672/")

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
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
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Client = conn
	Channel = ch
	Queue = q
	Context = ctx
	Cancel = cancel

	/*topic := "my-topic"
	partition := 0

	connection, err := kafka.DialLeader(context.Background(), "tcp", "192.168.0.213:9092", topic, partition)

	if err != nil {
		panic(err)
	}

	connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

	Client = connection*/
}

func Deinit() {
	Cancel()

	if err := Channel.Close(); err != nil {
		panic(err)
	}

	if err := Client.Close(); err != nil {
		panic(err)
	}
}
