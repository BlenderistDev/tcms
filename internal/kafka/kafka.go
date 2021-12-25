package kafka

import (
	"context"
	"fmt"
	"strings"
	"tcms/m/internal/dry"
	"time"

	kafka2 "github.com/segmentio/kafka-go"
)

type Client interface {
	Subscribe(ch chan []uint8)
}

type client struct {
	reader *kafka2.Reader
}

func NewKafkaClient() (Client, error) {
	kafkaURL, err := getKafkaHost()
	if err != nil {
		return nil, err
	}

	topic := "telegram-event"
	groupID := "tcms"

	brokers := strings.Split(kafkaURL, ",")
	reader := kafka2.NewReader(kafka2.ReaderConfig{
		Brokers:           brokers,
		GroupID:           groupID,
		Topic:             topic,
		MaxBytes:          10e6, // 10MB
		MaxWait:           time.Millisecond * 10,
		HeartbeatInterval: 1,
		ReadBackoffMax:    time.Millisecond * 100,
	})

	client := client{reader: reader}

	return client, nil
}

func (c client) Subscribe(ch chan []uint8) {
	fmt.Println("start consuming ... !!")
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			dry.HandleError(err)
		}
		ch <- m.Value
	}
}
