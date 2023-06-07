package models

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetESClient() (*elasticsearch.Client, string) {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://3.125.159.227:9200",
		},
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	indexName := "products"

	return es, indexName
}
