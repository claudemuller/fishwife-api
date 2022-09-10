package pkg

import (
	"log"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SendInternalServerError(w http.ResponseWriter) {
	log.Println("internal server error")
	w.WriteHeader(http.StatusInternalServerError)
}

func SendMethodNotSupported(w http.ResponseWriter) {
	log.Println("method not allowed")
	w.WriteHeader(http.StatusMethodNotAllowed)
}
