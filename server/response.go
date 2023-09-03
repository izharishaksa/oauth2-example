package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error any `json:"error"`
	Data  any `json:"data"`
}

func response(w http.ResponseWriter, data any, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{}
	if err != nil {
		resp.Error = err.Error()
	}
	if data != nil {
		resp.Data = data
	}

	w.WriteHeader(statusCode)
	jsonResponse, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResponse)
}
