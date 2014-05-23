package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"time"
)

const pollInterval = 10 * time.Second
const errTimeout = 10 * time.Second // back-off time
const statusInterval = 5 * time.Second
const numPollers = 3

type State struct {
	url    string
	status string
}

type Resource struct {
	url                    string
	errorsSinceLastSuccess int
}

func (r *Resource) Poll() string {
	res, err := http.Head(r.url)
	if nil != err {
		log.Println("Error", r.url, err)
		r.errorsSinceLastSuccess++
		return err.Error()
	}

	r.errorsSinceLastSuccess = 0
	return res.Status
}

func (r *Resource) Sleep(ready chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errorsSinceLastSuccess))
	ready <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func createResourcesFromUrls(filename string, out chan<- *Resource) {
	file, err := os.Open(filename)
	if nil != err {
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- &Resource{url: scanner.Text()}
	}
}

func logState(state map[string]string) {
	log.Println("Current state")
	for k, v := range state {
		log.Printf("%s %s", k, v)
	}
}

func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)

			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()

	return updates
}

func main() {
	pending, complete := make(chan *Resource), make(chan *Resource)

	status := StateMonitor(statusInterval)
	go createResourcesFromUrls("urls.txt", pending)
	// // uncommenting causes draining
	// for i := 0; i < 3; i++ {
	// 	x := <-pending
	// 	fmt.Println(x.url)
	// }

	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	for r := range complete {
		go r.Sleep(pending)
	}
}
