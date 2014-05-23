package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(url string) []byte {
	resp, err := http.Get(url)
	if nil != err {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		log.Fatal(err)
	}

	return body
}

type (
	youtubeVideo struct {
		Entry ytEntry `json:"entry"`
	}

	ytEntry struct {
		Title      ytTitle      `json:"title"`
		MediaGroup ytMediaGroup `json:"media$group"`
	}

	ytMediaGroup struct {
		MediaThumbnail []ytMediaThumbnail `json:"media$thumbnail"`
	}

	ytMediaThumbnail struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Name   string `json:"yt$name"`
	}

	ytTitle struct {
		Value string `json:"$t"`
	}
)

func main() {
	log.Println("Fetching...")
	data := fetch("http://gdata.youtube.com/feeds/api/videos/dyfFQw3RmBE?v=2&prettyprint=true&alt=json")
	log.Println("Done.")
	log.Println(string(data))
	var y youtubeVideo
	log.Println("Unmarshalling...")
	err := json.Unmarshal(data, &y)
	log.Println("Done.")
	if nil != err {
		log.Fatal(err)
	}

	log.Printf("%+v", y)
}
