package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Status int `json:status`
	Message string `json:message`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "BadRequest",
		})
		return 
	}

	fmt.Fprintf(w, "hello GO!")
}