package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	//"github.com/freshman-tech/news-demo-starter-files/news"
	//"github.com/freshman-tech/news-demo-starter-files/tides"
	"github.com/edinel/tides/news"
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
 *tideHandler deals with tide URLs.
 *Takes an input of a pointer to a tides.Client (defined in tides.go) and an http HandlerFunk
 *This uses a closure which means we're returning a function to the caller.  (I think)
 *
 */
func tideHandler(tidesapi *tides.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		tideInput := params.Get("T")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		results, err := tidesapi.FetchEverything(tideInput, page)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Tide Station is: ", tideInput)
		fmt.Println("Page is: ", page)
		fmt.Printf("%+v", results)
		//We need to do SOMETHING better than just dump these to stdout but maybe not tonight.

	}
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

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Search Query was: ", searchQuery)
		fmt.Println("Page is: ", page)
		fmt.Printf("%+v", results)

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
	tideapi := tides.NewClient(myTideClient, apiKey, 20)

	httpMux := http.NewServeMux()
	httpMux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	httpMux.HandleFunc("/", indexHandler)
	httpMux.HandleFunc("/search", searchHandler(newsapi))
	httpMux.HandleFunc("/tides", tideHandler(tideapi))
	http.ListenAndServe(":"+port, httpMux)
}

/*

https://api.tidesandcurrents.noaa.gov/mdapi/prod/webapi/stations/9414290.json?expand=details,sensors&units=metric
*/
