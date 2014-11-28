package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/ajstarks/svgo"
	rss "github.com/jteeuwen/go-pkg-rss"
)

var feeds []Feed = []Feed{
	mkFeed("http://reddit.com/r/dailyprogrammer/.rss"),
	mkFeed("http://www.echojs.com/rss"),
	mkFeed("http://www.infoq.com/feed?token=zqwZmgUydqtawZ3zoUvp2Q2xZcNoS7J5"),
}

var colors []string = []string{"#ff9999", "#99ff99", "#9999ff"}
var startTime int64 = now()
var timeout int
var feedSpace int
var wg sync.WaitGroup

func grabFeed(feed *Feed, feedChan chan bool, osvg *svg.SVG) {
	startGrab := time.Now().Unix()
	startGrabSeconds := startGrab - startTime

	fmt.Println("Grabbing feed", feed.url, "at", startGrabSeconds, "second mark")
	if FEED_NOT_READ == feed.status {
		fmt.Println("Feed not yet read")
		feed.status = FEED_DOWNLOADING

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
		endX := int((endSec - startGrab))
		if 0 == endX {
			endX = 1
		}

		endX *= 33

		fmt.Println("Read feed in", endX, "seconds")
		fmt.Println("startX:", startX, ", endX:", endX)
		osvg.Rect(startX, startY, endX, feedSpace, "fill:#000000;opacity:.4")
		wg.Wait()

		endGrab := time.Now().Unix()
		endGrabSeconds := endGrab - startTime
		feedEndX := int(endGrabSeconds * 33)

		osvg.Rect(feedEndX, startY, 1, feedSpace, "fill:#ff0000;opacity:.9;")

		feedChan <- true
	} else if FEED_DOWNLOADING == feed.status {
		fmt.Println("Feed already in progreess")
	}
}

func channelHandler(feed *rss.Feed, channels []*rss.Channel) {

}

func itemsHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Println("Found", len(newitems), "items in", feed.Url)

	for i := range newitems {
		url := newitems[i].Guid
		if nil == url {
			fmt.Println("nil")
		} else {
			fmt.Println(*url)
			_, err := http.Get(*url)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Done()
}

func getRSS(rw http.ResponseWriter, req *http.Request) {
	startTime = time.Now().Unix()
	rw.Header().Set("Content-Type", "image/svg+xml")
	outputSVG := svg.New(rw)
	outputSVG.Start(width, height)

	feedSpace = (height - 20) / len(feeds)

	for i := 0; i < 30000; i++ {
		timeText := strconv.FormatInt(int64(i/10), 10)
		if i%1000 == 0 {
			outputSVG.Text(i/30, 390, timeText, "text-anchor:middle;font-size:10px;fill:#000000")
		} else if i%4 == 0 {
			outputSVG.Circle(i, 377, 1, "fill:#cccccc;stroke:none")
		}

		if i%10 == 0 {
			outputSVG.Rect(i, 0, 1, 400, "fill:#dddddd")
		}
		if i%50 == 0 {
			outputSVG.Rect(i, 0, 1, 400, "fill:#cccccc")
		}
	}

	feedChan := make(chan bool, 3)

	for i := range feeds {
		outputSVG.Rect(0, (i * feedSpace), width, feedSpace, "fill:"+colors[i]+";stroke:none;")
		feeds[i].status = FEED_NOT_READ
		go grabFeed(&feeds[i], feedChan, outputSVG)
	}

	for _ = range feeds {
		<-feedChan
	}

	outputSVG.End()
}

func main() {
	runtime.GOMAXPROCS(2)

	timeout = 1000

	http.Handle("/getrss", http.HandlerFunc(getRSS))
	err := http.ListenAndServe(":1900", nil)
	if err != nil {
		panic(err)
	}
}
