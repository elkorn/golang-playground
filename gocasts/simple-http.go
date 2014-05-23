package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/repos/dotcloud/docker")
	if nil != err {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(body)

	if nil != err {
		log.Fatal(err)
	}
}
