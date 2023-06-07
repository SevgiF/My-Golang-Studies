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

func Filter(ctx *fiber.Ctx) error {

	name := ctx.FormValue("category")
	brands := ctx.FormValue("brands")
	pageStr := ctx.Query("p")

	brandArr := strings.Split(brands, ",")

	filter := []map[string]interface{}{}
	if name != "" {
		filter = append(filter, map[string]interface{}{"term": map[string]interface{}{"name": name}})
	}
	if len(brandArr) > 0 {
		filter = append(filter, map[string]interface{}{"terms": map[string]interface{}{"brand": brandArr}})
	}

	/* filter := make(map[string]interface{})
	if name != "" {
		filter["terms"] = map[string]interface{}{"name": name}
	}
	if len(brandArr) > 0 {
		filter["terms"] = map[string]interface{}{"brand": brandArr}
	} */

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
	}

	data, _ := json.Marshal(query)
	es, indexName := m.GetESClient()
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

	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write([]byte(res.String()))

	return nil
}
