package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	nString := query.Get("n")
	n, _ := strconv.Atoi(nString)
	ans := Fib(n)
	fmt.Fprint(w, ans)
}

func Fib(n int) int64 {
	if n == 1 || n == 2 {
		return 1
	} else {
		var a, b int64 = 1,1 
		for i:= 3; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
}