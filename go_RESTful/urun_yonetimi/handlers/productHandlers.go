package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	h "sevgifidan.com/urunYonetimi/helpers"
	m "sevgifidan.com/urunYonetimi/models"
)

// * Uygulamada veri tabanı kullanmayacağımız için,
// * veriyi kaydetme(uygulama hafızasında tutma), silme, güncelleme gibi işlemleri yapabilmek için,
// * bir nesneye ihtiyaç var. Bu nesne map nesnesi olarak seçtik.
var productStore = make(map[string]m.Product, 0)
var id int = 0

// todo: HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /api/products write postProduct
	//
	// Adding a product
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   schema:
	//     $ref: "#/definitions/Product"
	// responses:
	//   '200':
	//     description: product response

	var product m.Product
	err := json.NewDecoder(r.Body).Decode(&product) //requestin body'sinde bulunan veriyi decode edip product nesnesinin hafıza adresine aktar
	h.CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	h.CheckError(err)

	w.Header().Set("Content-Type", "application/json") //oluşturduğum verinin content tipini belirtiyorum
	w.WriteHeader(http.StatusCreated)                  //Yaptığımız oluşturma işlemi başarılı demek
	w.Write(data)
}

// todo: HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /api/products read getProducts
	//
	// Get all products
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: product response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Product"

	var products []m.Product
	for _, product := range productStore {
		products = append(products, product)
	}
	data, err := json.Marshal(products)
	h.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// todo: HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /api/products/{id} read getProduct
	//
	// Get a products
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: id
	//   required: true
	//   type: integer
	// responses:
	//   '200':
	//     description: product response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Product"

	var product m.Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	h.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// todo: HTTP Put - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /api/products/{id} write putProduct
	//
	// Updating a product
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: id
	//   required: true
	//   type: integer
	// - name: Body
	//   in: body
	//   schema:
	//     $ref: "#/definitions/Product"
	// responses:
	//   '200':
	//     description: product response

	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodUpd m.Product
	err = json.NewDecoder(r.Body).Decode(&prodUpd)
	h.CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUpd.ID, _ = strconv.Atoi(key)
		prodUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpd
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}

// todo: HTTP Delete - /api/product{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /api/products/{id} write deleteProduct
	//
	// Deleting a product
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - in: path
	//   name: id
	//   required: true
	//   type: integer
	// - in: body
	//   name: Body
	//   schema:
	//     $ref: "#/definitions/Product"
	// responses:
	//   '200':
	//     description: product response

	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
