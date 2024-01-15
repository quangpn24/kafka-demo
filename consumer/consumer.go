package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to consume messages
	topic := "helloo"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092", "localhost:9095", "localhost:9094"},
		Topic:   topic,
		//GroupID:        "my-group",
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		Partition:      1,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			println("err")
			break
		}
		fmt.Printf("message at offset %d: %s = %s. Partition: %v\n", m.Offset, string(m.Key), string(m.Value), m.Partition)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
