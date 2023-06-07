package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	m "github.com/SevgiF/elastic_kafka/models"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

func PostLog(ctx *fiber.Ctx) error {
	es, indexName := m.GetESClient()
	kafkaWriter := m.GetKafka()

	fields := []string{"name", "brand"}
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  "apple",
				"fields": fields,
			},
		},
	}

	data, _ := json.Marshal(esQuery)

	//indexs := []string{indexName}

	res, _ := es.Search(
		es.Search.WithContext(ctx.Context()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(bytes.NewBuffer(data)),
		/* es.Search.WithResponseListener(func(c context.Context, res *esapi.Response) {
			if res != nil {
				msg := kafka.Message{
					Key:   []byte("elastic"),
					Value: []byte(res.String()),
				}
				if err := kafkaWriter.WriteMessages(c, msg); err != nil {
					log.Printf("Error writing message to Kafka: %s", err)
				}
			}
		}), */
	)

	if res.IsError() {
		log.Fatalf("Error response received from Elasticsearch: %s", res.Status())
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write([]byte(res.String()))

	message := kafka.Message{
		Key:   []byte("first"),
		Value: data,
	}
	if err := kafkaWriter.WriteMessages(ctx.Context(), message); err != nil {
		log.Fatalf("Error writing message to Kafka: %s", err)
	}

	fmt.Println("Elasticsearch query sent to Kafka successfully.")

	return nil
}
