package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := Articles{
		Article{Title: "This is a test 1", Desc: "This is test desc", Content: "This is a test content"},
		Article{Title: "This is a test 2", Desc: "This is test desc", Content: "This is a test content"},
		Article{Title: "This is a test 3", Desc: "This is test desc", Content: "This is a test content"},
		Article{Title: "This is a test 4", Desc: "This is test desc", Content: "This is a test content"},
	}

	_ = json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
