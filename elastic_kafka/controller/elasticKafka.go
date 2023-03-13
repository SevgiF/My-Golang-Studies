package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	m "github.com/SevgiF/elastic_kafka/models"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

func Post(ctx *fiber.Ctx) error {
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

	indexs := []string{indexName}
	req := esapi.SearchRequest{Index: indexs, Body: strings.NewReader(string(data))}
	res, _ := req.Do(ctx.Context(), es)
	defer res.Body.Close()
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
