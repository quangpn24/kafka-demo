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
	topic := "helloo"

	w := &kafka.Writer{
		Addr:     kafka.TCP("127.0.0.1:9092", "127.0.0.1:9095", "127.0.0.1:9094"),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}

	for i := 0; i < 20; i++ {
		key := "MY KEY"
		message := fmt.Sprintf("Message %v", i)
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(key),
				Value: []byte(message),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	time.Sleep(5 * time.Second)
}
