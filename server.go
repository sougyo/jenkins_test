package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	filePath := "index.html"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err readfile")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(content))
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
