package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type articles []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	allArticles := articles{
		article{Title: "title one", Description: "description one"},
		article{Title: "title two", Description: "description two"},
		article{Title: "title three", Description: "description three"},
		article{Title: "title four", Description: "description four"},
		article{Title: "title five", Description: "description five"},
	}

	fmt.Println("end points hit all the articles")
	json.NewEncoder(w).Encode(allArticles)
}

func homHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page endpoint hits")
}

func requestHandler() {
	http.HandleFunc("/", homHandler)
	http.HandleFunc("/articles", allArticles)
}

func main() {
	requestHandler()
}