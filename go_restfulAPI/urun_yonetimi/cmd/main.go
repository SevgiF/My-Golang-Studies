package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	h "sevgifidan.com/urunYonetimi/handlers"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter().StrictSlash(true) //url'in sonuna / gelirse de sayfayı gösterir (örn: /api/products/)
	r.HandleFunc("/api/products", h.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", h.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", h.PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", h.PutProductHandler).Methods("Put")
	r.HandleFunc("/api/products/{id}", h.DeleteProductHandler).Methods("Delete")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	log.Println("Server ending...")
}
