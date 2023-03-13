package models

import (
	"github.com/segmentio/kafka-go"
)

func GetKafka() *kafka.Writer {
	conn := kafka.WriterConfig{
		Brokers: []string{"3.69.23.63:9092"},
		Topic:   "elastic",
	}

	kafkaWriter := kafka.NewWriter(conn)

	return kafkaWriter
}
