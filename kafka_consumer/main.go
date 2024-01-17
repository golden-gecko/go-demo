package main

import (
    "context"
    "log"

    kafka "github.com/segmentio/kafka-go"
)

func main() {
    topic := "my-topic"
    partition := 0

    r := kafka.NewReader(
        kafka.ReaderConfig{
            Brokers:   []string{"192.168.0.213:9092"},
            Topic:     topic,
            Partition: partition,
            MinBytes:  10e3, // 10KB
            MaxBytes:  10e6, // 10MB
        },
    )

    defer r.Close()

    for {
        m, err := r.ReadMessage(context.Background())

        if err != nil {
            break
        }

        log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }
}
