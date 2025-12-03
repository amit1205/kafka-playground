package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/amit1205/kafka-playground/internal/kafka"
	"github.com/amit1205/kafka-playground/internal/model"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{kafka.BootstrapServers},
		Topic:          kafka.OrderTopic,
		GroupID:        kafka.OrderGroupID,
		CommitInterval: time.Second, // commit offsets every second
		MinBytes:       1,
		MaxBytes:       10e6,
	})

	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("failed to close reader: %v", err)
		}
	}()

	log.Printf("Starting order consumer. Topic=%q GroupID=%q",
		kafka.OrderTopic, kafka.OrderGroupID)

	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v", err)
			// In a real service you’d handle retries / backoff / shutdown.
			continue
		}

		var order model.Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("failed to unmarshal order: %v", err)
			continue
		}

		processOrder(order)
	}
}

func processOrder(o model.Order) {
	// Simulate doing something meaningful with the order.
	log.Printf("[PROCESSOR] Received order: ID=%s User=%s Amount=%.2f %s Status=%s",
		o.ID, o.UserID, o.Amount, o.Currency, o.Status)

	// Here you could:
	//   - Charge payment
	//   - Update inventory
	//   - Publish another Kafka event: OrderPaid, OrderFailed, etc.
	//   - Update a DB
	time.Sleep(200 * time.Millisecond)
}
