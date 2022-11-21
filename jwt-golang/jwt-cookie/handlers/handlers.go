package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	m "sevgifidan.com/jwtGolang/jwt-cookie/models"
)

// *key used to create signature
var key = []byte("super_secret_key")

// *valid users
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// todo: Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds m.Credentials
	//* Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		//if the structure of the body is wrong, return HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//* Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	//* If a password exists for the given user
	//* AND, if its same as password we received, the we can continue
	//* IF NOT, then we return "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//*Declare the expiration time of the token
	//here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(1 * time.Minute)
	iatTime := time.Now()
	//* Create the JWT claims, which includes the username and expiry time
	claims := &m.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  iatTime.Unix(),
		},
	}

	//* Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//* create JWT string
	tokenStr, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//* Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set expiry time  which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expirationTime,
	})
}

// todo: Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	// we can obtain(take) the session token from the requests cookies, which come every requests
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//* Get the JWT string from the cookie
	tknStr := c.Value

	//* Initialize a new instance of "Claims"
	claims := &m.Claims{}

	//* Parse the JWT string and store the result in "claims"
	//* This method will return an error if the token is invalid
	//(if it has expired according to the expiry time we set on sign in or if the signature does not match)
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Finally return the welcome message to the user, along with their username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

// todo: Refresh handler
func Resfresh(w http.ResponseWriter, r *http.Request) {
	//BEGIN (the code uptil this point is the same as the first part of the Welcome)
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &m.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//END

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within 30 seconds of expiry.
	// Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//* Now create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(1 * time.Minute)
	iatTime := time.Now()
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = iatTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//* Set the new token as the users "token" cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expirationTime,
	})
}
