package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var users = map[string]string{
	"username": "password",
}

var clients = map[string]Client{
	"client_id": {
		Id:          "client_id",
		Secret:      "client_secret",
		CallbackUrl: "http://localhost:9191/callback",
	},
}

var authCodeToClientID = map[string]string{}

var secretKey = []byte("your-secret-key")

type Client struct {
	Id          string
	Secret      string
	CallbackUrl string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", authenticationPage).Methods("GET")
	r.HandleFunc("/authorize", authorizationPage).Methods("GET")
	r.HandleFunc("/authenticate", authenticate).Methods("POST")
	r.HandleFunc("/authorize", authorize).Methods("POST")
	r.HandleFunc("/token", getAccessToken).Methods("POST")

	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
