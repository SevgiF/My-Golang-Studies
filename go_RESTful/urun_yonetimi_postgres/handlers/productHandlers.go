package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	h "sevgifidan.com/urunYonetimi/helpers"
	m "sevgifidan.com/urunYonetimi/models"
)

// todo: Veritabanı bağlantısı
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "9m61t1A-"
	dbname   = "productmanagement"
)

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode = disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connString)
	h.CheckError(err)
}

// todo: HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product m.Product
	err := json.NewDecoder(r.Body).Decode(&product) //requestin body'sinde bulunan veriyi decode edip product nesnesinin hafıza adresine aktar
	h.CheckError(err)

	product.CreatedOn = time.Now()
	result, err := db.Exec("INSERT INTO product(prdname, description, createdon) VALUES($1, $2, $3)", product.Name, product.Description, product.CreatedOn)
	h.CheckError(err)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)\n", rowsAffected)

	data, err := json.Marshal(product)
	h.CheckError(err)

	w.Header().Set("Content-Type", "application/json") //oluşturduğum verinin content tipini belirtiyorum
	w.WriteHeader(http.StatusCreated)                  //Yaptığımız oluşturma işlemi başarılı demek
	w.Write(data)
}

// todo: HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []*m.Product
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		prd := &m.Product{}
		err := rows.Scan(&prd.ID, &prd.Name, &prd.Description, &prd.CreatedOn, &prd.ChangedOn)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	data, err := json.Marshal(products)
	h.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// todo: HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var prd m.Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	err := db.QueryRow("SELECT * FROM product WHERE id=$1", key).Scan(&prd.ID, &prd.Name, &prd.Description, &prd.CreatedOn, &prd.ChangedOn)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product: %d %s-%s\n", prd.ID, prd.Name, prd.Description)
	}

	data, err := json.Marshal(prd)
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

	prodUpd.ChangedOn = time.Now()
	result, err := db.Exec("UPDATE product SET prdname=$1, description=$2, changedon=$3 WHERE id=$4", prodUpd.Name, prodUpd.Description, prodUpd.ChangedOn, key)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)\n", rowsAffected)

	w.WriteHeader(http.StatusOK)
}

// todo: HTTP Delete - /api/product{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	result, err := db.Exec("DELETE FROM product WHERE id=$1", key)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)\n", rowsAffected)

	w.WriteHeader(http.StatusOK)
}
