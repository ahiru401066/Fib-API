package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"fib_api/internal/service"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	nString := query.Get("n")
	n, _ := strconv.Atoi(nString)
	ans := service.Fib(n)
	fmt.Fprint(w, ans)
}