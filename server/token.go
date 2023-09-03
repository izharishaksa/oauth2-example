package main

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"time"
)

func validateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	return false, nil
}

func generateAuthToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateAuthCode(clientID string) (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	authCode := base64.URLEncoding.EncodeToString(randomBytes)

	return authCode, nil
}

func getAccessToken(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")
	clientId, err := validateAuthCode(authCode)
	if err != nil {
		response(w, nil, err, http.StatusUnauthorized)
		return
	}

	accessToken, err := generateAccessToken(clientId)
	if err != nil {
		response(w, nil, err, http.StatusInternalServerError)
		return
	}

	response(w, accessToken, nil, http.StatusOK)
}

func validateAuthCode(authCode string) (string, error) {
	clientID, ok := authCodeToClientID[authCode]
	if !ok {
		return "", fmt.Errorf("invalid auth_code")
	}

	return clientID, nil
}

func generateAccessToken(clientID string) (string, error) {
	accessToken := "dummy_access_token"

	return accessToken, nil
}
