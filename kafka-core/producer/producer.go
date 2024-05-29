package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to produce messages
	topic := "ORDER_CREATED_TOPIC"

	w := &kafka.Writer{
		Addr:     kafka.TCP("127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}

	key := "MY KEY"
	message := fmt.Sprintf("Message")
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	time.Sleep(5 * time.Second)
}
