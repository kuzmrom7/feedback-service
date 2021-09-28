package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (r *ResponseError) Respond(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}

func (r *ResponseReviews) SuccessRespond(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}
