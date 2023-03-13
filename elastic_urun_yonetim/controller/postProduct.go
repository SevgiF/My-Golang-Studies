package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	m "github.com/sfidann/urun_yonetim_elastic/models"
	u "github.com/sfidann/urun_yonetim_elastic/utils"
)

func PostProduct(ctx *fiber.Ctx) error {
	product := &m.Product{}
	err := json.Unmarshal(ctx.Body(), product)
	if err != nil {
		res := u.Response{Success: false, Message: "Invalid request"}
		data, _ := json.Marshal(res)
		ctx.Write(data)
	}

	//connection database
	conn := m.GetDatabase()
	defer m.CloseDatabase(conn)

	_, err = conn.Exec("INSERT INTO products(id, name, brand, price, quantity) VALUES($1,$2,$3,$4,$5)", product.ID, product.Name, product.Brand, product.Price, product.Quantity)
	if err != nil {
		res := u.Response{Success: false, Message: "Product adding failed."}
		data, _ := json.Marshal(res)
		ctx.Write(data)
	}

	data, _ := json.Marshal(product)
	dID := strconv.Itoa(product.ID)
	//create Elasticsearch index
	es, indexName := m.GetESClient()

	req := esapi.IndexRequest{Index: indexName, DocumentID: dID, Body: strings.NewReader(string(data))}
	res, _ := req.Do(ctx.Context(), es)
	fmt.Println(res)

	fmt.Printf("Indexed product with ID %d to index %s\n", product.ID, indexName)
	resp := u.Response{Success: true, Message: "Product added.", Data: product}
	response, _ := json.Marshal(resp)
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(response)

	return nil
}
