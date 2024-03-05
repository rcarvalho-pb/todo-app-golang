package response_json

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(struct {
			Status  int `json:"status,omitempty"`
			Content any `json:"content,omitempty"`
		}{
			Status:  statusCode,
			Content: data,
		}); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(struct {
			Status  int
			Content any
		}{
			Status:  statusCode,
			Content: "no content",
		}); err != nil {
			log.Fatal(err)
		}
	}
}
