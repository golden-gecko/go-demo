package queue

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
    Client  *amqp.Connection
    Channel *amqp.Channel
    Queue   amqp.Queue
    Context context.Context
    Cancel  context.CancelFunc

    KafkaClient *kafka.Conn
)

func Init() error {
    conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")

    if err != nil {
        return err
    }

    ch, err := conn.Channel()

    if err != nil {
        return err
    }

    err = ch.Qos(10, 0, true)

    queues := []string{"temperatures", "transits"}

    for _, queueName := range queues {
        _, err := ch.QueueDeclare(
            queueName, // name
            false,     // durable
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

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

    Client = conn
    Channel = ch
    Context = ctx
    Cancel = cancel

    return nil
}

func Deinit() error {
    Cancel()

    if err := Channel.Close(); err != nil {
        return err
    }

    if err := Client.Close(); err != nil {
        return err
    }

    return nil
}

func InitKafka() error {
    topic := "my-topic"
    partition := 0

    connection, err := kafka.DialLeader(context.Background(), "tcp", "192.168.0.213:9092", topic, partition)

    if err != nil {
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
        return err
    }

    return nil
}
