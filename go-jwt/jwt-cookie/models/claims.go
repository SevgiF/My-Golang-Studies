package models

import "github.com/dgrijalva/jwt-go"

//todo: struct that will be encoded to JWT
//* StandartClaims: to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
