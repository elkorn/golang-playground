package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func First(query string, replicas ...Search) Result {
	// buffered channels makes the goroutines exit on their own, not leaving any garbage behind.
	c := make(chan Result, len(replicas))
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result, 3)
	// go func() { c <- Web(query) }()
	// go func() { c <- Video(query) }()
	// go func() { c <- Image(query) }()
	go func() { c <- First(query, Web, Web) }()
	go func() { c <- First(query, Video, Video) }()
	go func() { c <- First(query, Image, Image) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < cap(c); i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s results for %s\n", kind, query))
	}
}

func main() {
	fmt.Println(Google("test"))
}
