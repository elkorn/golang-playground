package main

import (
	"fmt"
	"time"
)

// godoc.org -> a generated documentation for Go packages.

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}

	go s.loop()
	return s
}

type naiveSub struct {
	fetcher Fetcher
	updates chan Item
}

type sub struct {
	closing chan chan error // channel on which you pass error channels
	naiveSub
}

func (s *naiveSub) loop() {
	for {
		// The BUG1s can be detected on a 64-bit system by running Go with the `-race` parameter.
		if s.closed { // BUG1: unsynchronized access to s.closed
			close(s.updates)
			return
		}

		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err                  // BUG1: unsynchronized access to s.err
			time.Sleep(10 * time.Second) // BUG2: may keep loop running if updates are rare.
			continue
		}

		for _, item := range items {
			s.updates <- item //BUG3: This will block forever if nobody will be able to receive the update.
		}

		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now)) // BUG3: may keep loop running if updates are rare.
		}
	}
}

func (s *naiveSub) Updates() <-chan Item {
	return s.updates
}

func (s *naiveSub) Close() error {
	s.closed = true // BUG1: unsynchronized access to s.closed
	return s.err    // BUG1: unsynchronized access to s.err
}

// the set of pending items may grow without bounds if the downstream receiver is busy -> backpressure
const maxPending = 10

func Fetch(url string) {
	var pending []Item
	var next time.Time             // initally Jan 1, year 0
	var fetchDone chan fetchResult // if non-nil. a Fetch is running (guard for slow web resource-related blocking)
	var err error

	for {
		var fetchDelay time.Duration // when will the next fetch occur; initally 0 (no delay)
		if now := time.Now; next.After(now) {
			// If we're not supposed to fetch right now, calculate the delay after which it should be done.
			fetchDelay = next.Sub(now)
		}

		var startFetch <-chan time.Time
		// 1) Don't block for slow resources - only fetch if a previous fetch has returned. (instead of rescheduling)
		// 2) Handle backpressure - fetch only if there are less than N pending items to avoid hogging memory.
		if fetchDone == nil && len(pending) < maxPending {
			startFetch := time.After(fetchDelay)
		}

		select {
		case <-startFetch:
			var fetched []Item
			fetchDone = make(chan fetchResult, 1)
			go func() {
				fetched, next, err = s.fetcher.Fetch()
				fetchedDone <- fetchResult{fetched, next, err}
			}()

		case result := <-fetchDone:
			fetchDone := nil
			// Append all the items to the pending queue.
			pending = append(pending, fetched...)
		}

	}
}

func Send() {
	var pending []Item
	for {
		select {
		case s.updates <- pending[0]:
			pending = pending[1:]
		}
	}
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

type fetchResult struct {
	fetched []Item
	next    time.Time
	err     error
}

// A for/select loop prevents data races.
func (s *sub) loop() {
	// ... declare mutable state ...
	for {
		var pending []Item
		var next time.Time
		var err error
		seen := make(map[string]bool)
		// ... set up channels for cases ...
		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates) // Tells the receiver that we're done.
			return

		case <-startFetch:
			var fetched []Item
			fetched, next, err = s.fetcher.Fetch() // may return duplicates.
			// fetcher.Fetch() might block for slow web resources.
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}

			for _, item := range fetched {
				if !seen[item.GUID] {
					pending = append(pending, item)
					seend[item.GUID] = true
				}
			}

		case updates <- pending[0]:
			pending = pending[1:]
		}
	}
}

func main() {
	merged := Merge(
		Subscribe(Fetch("blog.golang.com")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")),
	)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}
}
