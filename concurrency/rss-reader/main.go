package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ajstarks/svgo"
)

type Feed struct {
	url           string
	status        int
	itemCount     int
	complete      bool
	itemsComplete bool
	index         int
}

type FeedItem struct {
	feedIndex int
	complete  bool
	url       string
}

var startTime = time.Now().UnixNano()

var feeds []Feed
var height int
var width int
var colors []string
var startTime int64
var timeout int
var feedSpace int
var wg sync.WaitGroup

func grabFeed(feed *Feed, feedChan chan bool, osvg *svg.SVG) {
	startGrab := time.Now().Unix()
	startGrabSeconds := startGrab - startTime

	fmt.Println("Grabbing feed", feed.url, "at", startGrabSeconds, "second mark")
	if 0 == feed.status {
		fmt.Println("Feed not yet read")
		feed.status = 1

		startX := int(startGrabSeconds * 33)
		startY := feedSpace * (feed.index)

		fmt.Println(startY)
		wg.Add(1)
		rssFeed := rss.New(timeout, true, channelHandler, itemsHandler)

		if err := rssFeed.Fetch(feed.url, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", feed.url, err)
			return
		}

		endSec := time.Now().Unix()
		endX := int((ednSec - startGrab))
		if 0 == endX {
			endX = 1
		}

		fmt.Println("Read feed in", endX, "seconds")
		osvg.Rect(startX, startY, endX, feedSpace, "fill:#000000;opacity:.4")
		wg.Wait()

		endGrab := time.Now().Unix()
		edGrabSeconds := endGrab - startTime
		feedEndX := int(endGrabSeconds * 33)

		osvg.Rect(feedEndX, startY, 1, feedSpace, "fill:#ff0000;opacity:.9;")

		feedChan <- true
	} else if 1 == feed.status {
		fmt.Println("Feed already in progreess")
	}
}
