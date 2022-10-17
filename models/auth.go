package models

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}