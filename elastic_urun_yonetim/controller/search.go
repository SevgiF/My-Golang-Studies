package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	m "github.com/sfidann/urun_yonetim_elastic/models"
)

func Search(ctx *fiber.Ctx) error {

	query := ctx.Query("q")
	pageStr := ctx.Query("p")

	fields := []string{"name", "brand"}
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": fields,
			},
		},
	}

	data, _ := json.Marshal(esQuery)
	es, indexName := m.GetESClient()
	/* req, err := es.Search(es.Search.WithContext(ctx.Context()), es.Search.WithIndex(indexName), es.Search.WithBody(strings.NewReader(string(data))))
	res := req.String()
	if err != nil {
		log.Fatalf("Error executing search query: %s", err)
	} */

	//* Pagination
	size := 5
	page, _ := strconv.Atoi(pageStr)
	started_record := page * size

	indexs := []string{indexName}
	req := esapi.SearchRequest{Index: indexs, Body: strings.NewReader(string(data)), From: esapi.IntPtr(started_record), Size: esapi.IntPtr(size)} //From: esapi.IntPtr(2), Size: esapi.IntPtr(3)

	res, _ := req.Do(ctx.Context(), es)

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response received from Elasticsearch: %s", res.Status())
	}
	//fmt.Println(res)

	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write([]byte(res.String()))

	return nil

}
