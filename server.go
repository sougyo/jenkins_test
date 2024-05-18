package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sougyo/jenkins_test/goban"
)

func parseMove(input string) (rune, int, int, error) {
	parts := strings.Fields(input)
	if len(parts) != 4 {
		return goban.Empty, 0, 0, fmt.Errorf("invalid input format")
	}

	stone := goban.Empty
	s := parts[1]
	if s == "WHITE" {
		stone = goban.White
	} else if s == "BLACK" {
		stone = goban.Black
	} else {
		return goban.Empty, 0, 0, fmt.Errorf("failed to parse stone")
	}

	n1, err := strconv.Atoi(parts[2])
	if err != nil {
		return goban.Empty, 0, 0, fmt.Errorf("failed to parse number 1: %v", err)
	}

	n2, err := strconv.Atoi(parts[3])
	if err != nil {
		return goban.Empty, 0, 0, fmt.Errorf("failed to parse number 2: %v", err)
	}

	return stone, n1, n2, nil
}

func main() {
	board := goban.NewGoBoard()

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
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		stone, n1, n2, err := parseMove(string(input))
		if err != nil {
			http.Error(w, "Failed to read request body(unexpected format)", http.StatusInternalServerError)
			return
		}

		board.PlaceStone(n1, n2, stone)

		_, err = w.Write([]byte("OK"))
		if err != nil {
			http.Error(w, "Failed to write request body", http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte(board.Display()))
		if err != nil {
			http.Error(w, "Failed to write request body", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
