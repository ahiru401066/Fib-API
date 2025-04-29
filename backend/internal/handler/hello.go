package handler

import (
	"fmt"
	"net/http"

	"fib_api/internal/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}

	fmt.Fprintf(w, "hello GO!")
}