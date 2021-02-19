package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func createArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(1000000))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(&article)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:_index], articles[_index+1:]...)

			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
			articles = append(articles, article)
			json.NewEncoder(w).Encode(&article)

			return
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Article Library")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", createArticles).Methods("POST")
	myRouter.HandleFunc("/articles/{id:[0-9]+}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id:[0-9]+}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id:[0-9]+}", deleteArticle).Methods("DELETE")
	http.ListenAndServe(":9000", myRouter)
}

func main() {

	articles = append(articles,
		Article{ID: "1", Title: "Test Title 1", Description: "Test Discription 1", Content: "Hello world 1"},
		Article{ID: "2", Title: "Test Title 2", Description: "Test Discription 2", Content: "Hello world 2"},
		Article{ID: "3", Title: "Test Title 3", Description: "Test Discription 3", Content: "Hello world 3"},
		Article{ID: "4", Title: "Test Title 4", Description: "Test Discription 4", Content: "Hello world 4"})

	handleRequest()
}
