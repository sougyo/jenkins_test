package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	filePath := "index.html"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err readfile")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write(content)
		if err != nil {
			log.Fatalf("err writing response: %v", err)
		}
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
