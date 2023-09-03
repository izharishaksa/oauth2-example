package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var clientId = "client_id"
var clientSecret = "client_secret"
var callbackUrl = "http://localhost:9191/callback"
var tokenUrl = "http://localhost:9090/token"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/callback", callback).Methods("GET")

	err := http.ListenAndServe(":9191", router)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<a href='http://localhost:9090/authorize?client_id=%s&callback=%s'>Sign Up Using Kancyl</a>", clientId, url.QueryEscape(callbackUrl))
}

func callback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")

	if code == "" {
		http.Error(w, "Authorization code missing", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", callbackUrl)

	resp, err := http.Post(tokenUrl, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error exchanging authorization code: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error exchanging authorization code: %v", resp.Status), resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, copyErr := io.Copy(w, resp.Body)
	if copyErr != nil {
		http.Error(w, fmt.Sprintf("Error copying response body: %v", copyErr), http.StatusInternalServerError)
		return
	}
}
