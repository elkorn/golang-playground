package main

var feedIndex = 0

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

const (
	FEED_NOT_READ = iota
	FEED_DOWNLOADING
)

func mkFeed(url string) (result Feed) {
	result = Feed{
		index:         feedIndex,
		url:           url,
		status:        FEED_NOT_READ,
		itemCount:     0,
		itemsComplete: false,
		complete:      false,
	}

	feedIndex++
	return result
}
