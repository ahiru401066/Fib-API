package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello GO!")
}

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

func main() {
	http.HandleFunc("/fib", FibHandler)
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server is running on :8080...")
	if err:= http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}