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
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
