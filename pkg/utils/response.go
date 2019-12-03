package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithError(err error) *Response {
	r.Error = err
	return r
}

func New(message string, status bool) *Response {
	return &Response{Status: status, Message: message}
}

func (r *Response) Respond(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}
