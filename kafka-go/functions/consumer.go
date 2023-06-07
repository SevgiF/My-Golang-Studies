package functions

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Consumer(brokers []string) {
	config := kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     "my-topic",
		Partition: 0,
		MinBytes:  10e3, // en az 10KB oku
		MaxBytes:  10e6, // en fazla 10MB oku
	}

	r := kafka.NewReader(config)
	defer r.Close()

	//*Mesajları oku
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("key=%s value=%s\n", string(m.Key), string(m.Value))
	}
}
