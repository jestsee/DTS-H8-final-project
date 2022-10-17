package models

import "github.com/dgrijalva/jwt-go"

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
