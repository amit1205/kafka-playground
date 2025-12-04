package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	kafkago "github.com/segmentio/kafka-go"

	config "github.com/amit1205/kafka-playground/internal/kafka"
	"github.com/amit1205/kafka-playground/internal/model"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	writer := &kafkago.Writer{
		Addr:         kafkago.TCP(config.BootstrapServers),
		Topic:        config.OrderTopic,
		Balancer:     &kafkago.LeastBytes{},
		RequiredAcks: kafkago.RequireAll,
		Async:        false,
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}()

	log.Printf("Starting order producer. Writing to topic %q", config.OrderTopic)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		order := randomOrder()
		payload, err := json.Marshal(order)
		if err != nil {
			log.Printf("failed to marshal order: %v", err)
			continue
		}

		msg := kafkago.Message{
			Key:   []byte(order.ID),
			Value: payload,
		}

		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			log.Printf("failed to write message: %v", err)
			continue
		}

		log.Printf("Produced order: ID=%s User=%s Amount=%.2f %s",
			order.ID, order.UserID, order.Amount, order.Currency)
	}
}

func randomOrder() model.Order {
	return model.Order{
		ID:        fmt.Sprintf("order-%d", rand.Intn(1000000)),
		UserID:    fmt.Sprintf("user-%d", rand.Intn(100)),
		Amount:    10 + rand.Float64()*90,
		Currency:  "USD",
		CreatedAt: time.Now().UTC(),
		Status:    "CREATED",
	}
}
