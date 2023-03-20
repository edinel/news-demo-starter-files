package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"net/url"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

/*
indexHandler tells you want to do with an index.
*/func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

/*
searchHandler desls with searches.  Unclear how it gets passed in?

*/
func searchHandler(w http.ResponseWriter, r *http.Request) {
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


/*
tideHandler deals with tide URLs
*/
func tideHandler(w http.ResponseWriter, r *httpRequest){
	u, err:url.Parse(r.Url.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := u.Query()
	SearchQuery:= params.Get("q")
	page := params.Get("Page")
	if page ==""{
		page = "1"
	}

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)


}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .evn file")
	}
	fs := http.FileServer(http.Dir("assets"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	httpMux.HandleFunc("/", indexHandler)
	httpMux.HandleFunc("/search", searchHandler)
	httpMux.HandleFunc("/tides", tideHandler)
	http.ListenAndServe(":"+port, httpMux)
}


/*

https://api.tidesandcurrents.noaa.gov/mdapi/prod/webapi/stations/9414290.json?expand=details,sensors&units=metric
*/
