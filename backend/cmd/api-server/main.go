package main

import (
	"fmt"
	"log"
	"net/http"

	"fib_api/internal/handler"
)

func main() {
	http.HandleFunc("/fib", handler.FibHandler)
	http.HandleFunc("/", handler.HelloHandler)

	fmt.Println("Server is running on :8080...")
	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}