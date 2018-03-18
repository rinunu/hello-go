package rss

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"log"
)

func MakeHandler(r *mux.Router, service *RssService) {
	r.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		AddSource(w, r, service)
	}).Methods("POST")

	r.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		GetAllSources(w, r, service)
	}).Methods("GET")

	r.HandleFunc("/sources/{sourceId}/articles", func(w http.ResponseWriter, r *http.Request) {
		GetAllArticles(w, r, service)
	}).Methods("GET")
}

type AddSourceParams struct {
	Url string `json:"url"`
}

func AddSource(w http.ResponseWriter, r *http.Request, service *RssService) {
	var post AddSourceParams

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Print(err)
		http.Error(w, "", 500)
	}

	println(post.Url)

	_, err = service.Subscribe(post.Url)
	if err != nil {
		log.Print(err)
		http.Error(w, "", 500)
	}
}

func GetAllSources(w http.ResponseWriter, r *http.Request, service *RssService) {
	sources := service.FindAllSources()

	json.NewEncoder(w).Encode(sources)
}

func GetAllArticles(w http.ResponseWriter, r *http.Request, service *RssService) {
	vars := mux.Vars(r)

	sourceId, err := strconv.Atoi(vars["sourceId"])
	if err != nil {
		panic(err)
	}

	articles, err := service.FindArticles(sourceId)
	if err != nil {
		// TODO
	}

	json.NewEncoder(w).Encode(articles)
}
