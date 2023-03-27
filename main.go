package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/freshman-tech/news-demo-starter-files/news"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

/*
indexHandler tells you want to do with an index.
*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

/*
tideHandler deals with tide URLs
*/
func tideHandler(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := url.Query()
	tideQuery := params.Get("T")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	fmt.Println("Tide ID is: ", tideQuery)
	fmt.Println("Page is: ", page)
}

/*
searchHandler desls with searches.  Unclear how it gets passed in?
*/
func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		fmt.Println("Search Query is: ", searchQuery)
		fmt.Println("Page is: ", page)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	fs := http.FileServer(http.Dir("assets"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myNewsClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myNewsClient, apiKey, 20)
	myTideClient := &http.Client{Timeout: 10 * time.Second}
	tideapi := tide.NewClient{myTideClient, apiKey, 20}

	httpMux := http.NewServeMux()
	httpMux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	httpMux.HandleFunc("/", indexHandler)
	httpMux.HandleFunc("/search", searchHandler(newsapi))
	httpMux.HandleFunc("/tides", tideHandler)
	http.ListenAndServe(":"+port, httpMux)
}

/*

https://api.tidesandcurrents.noaa.gov/mdapi/prod/webapi/stations/9414290.json?expand=details,sensors&units=metric
*/
