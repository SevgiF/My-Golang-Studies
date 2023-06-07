package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	h "sevgifidan.com/jwtGolang/jwt-cookie/handlers"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/signin", h.Signin)
	r.HandleFunc("/welcome", h.Welcome)
	r.HandleFunc("/refresh", h.Resfresh)

	log.Fatal(http.ListenAndServe(":8000", r))
}
