package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
	"os"
)

/**
 * [GetAll description]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func GetAll(w http.ResponseWriter, r *http.Request) {

	rss := map[string]interface{}{
		"Parser": map[string]string{
			"Google": "http://localhost:8080/news?feed=https://news.google.com/news/rss/?ned=pt-BR_br&gl=BR&hl=pt-BR",
			"msn":    "http://localhost:8080/news?feed=https://rss.msn.com/pt-br/",
			"Yahoo":  "http://localhost:8080/news?feed=http://rss.news.yahoo.com/rss/entertainment",
		},
	}
	json.NewEncoder(w).Encode(rss)
}

/**
 * [GetFeed description]
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func GetFeed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	feed := RSSfeed(params["feed"])
	json.NewEncoder(w).Encode(feed)
}

/**
 * [RSSfeed description]
 * @param {[type]} params string) (feed *gofeed.Feed [description]
 */
func RSSfeed(params string) (feed *gofeed.Feed) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return feed
}

/**
 * [main description]
 * @return {[type]} [description]
 */
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetAll).Methods("GET")
	router.HandleFunc("/news", GetFeed).Queries("feed", "{feed}").Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
