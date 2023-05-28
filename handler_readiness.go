package main

import (
	"net/http"
	"time"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handlerDay(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct {
		Day   int    `json:"day"`
		Month string `json:"month"`
	}{
		Day:   time.Now().Day(),
		Month: time.Now().Month().String(),
	})
}
