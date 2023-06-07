package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var key = []byte("cok-gizli-sifrem:)")

type Payload struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetJWT(w http.ResponseWriter, r *http.Request) {

	iat := 1664882773

	payload := &Payload{
		Username: "sevgi",
		StandardClaims: jwt.StandardClaims{
			IssuedAt: int64(iat),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, tokenStr)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/token", GetJWT).Methods("GET")
	http.ListenAndServe(":9000", router)
}
