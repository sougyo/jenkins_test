package main

import (
	"fmt"
	"net/http"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!!!")
	})
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
