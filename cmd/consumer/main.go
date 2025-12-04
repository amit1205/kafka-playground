package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	kafkago "github.com/segmentio/kafka-go"

	config "github.com/amit1205/kafka-playground/internal/kafka"
	"github.com/amit1205/kafka-playground/internal/model"
)

func main() {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:        []string{config.BootstrapServers},
		Topic:          config.OrderTopic,
		GroupID:        config.OrderGroupID,
		CommitInterval: time.Second,
		MinBytes:       1,
		MaxBytes:       10e6,
	})

	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("failed to close reader: %v", err)
		}
	}()

	log.Printf("Starting order consumer. Topic=%q GroupID=%q",
		config.OrderTopic, config.OrderGroupID)

	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v", err)
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
	log.Printf("[PROCESSOR] Received order: ID=%s User=%s Amount=%.2f %s Status=%s",
		o.ID, o.UserID, o.Amount, o.Currency, o.Status)

	time.Sleep(200 * time.Millisecond)
}
