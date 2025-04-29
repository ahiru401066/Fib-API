package utils

import(
	"net/http"
	"encoding/json"

	"fib_api/internal/model"
)

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	// json形式でレスポンス
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Status: http.StatusBadRequest,
		Message: message,
	})
} 