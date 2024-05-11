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
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		_, err = w.Write([]byte("OK"))
		if err != nil {
			http.Error(w, "Failed to write request body", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
