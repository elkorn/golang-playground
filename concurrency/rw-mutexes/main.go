package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeStruct struct {
	totalChanges int
	currentTime  time.Time
	rwLock       sync.RWMutex
}

func (self *TimeStruct) update() {
	self.rwLock.Lock()
	self.currentTime = time.Now()
	self.totalChanges++
	self.rwLock.Unlock()
}

var timeElement TimeStruct

func updateTime() {
	timeElement.update()
}

func main() {
	var wg sync.WaitGroup
	timeElement.totalChanges = 0
	timeElement.currentTime = time.Now()
	timer := time.NewTicker(time.Second)
	writeTimer := time.NewTicker(3 * time.Second)
	endTimer := make(chan bool)

	wg.Add(3)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println(
					timeElement.totalChanges,
					timeElement.currentTime.String())
			case <-writeTimer.C:
				updateTime()
				wg.Done()
			case <-endTimer:
				timer.Stop()
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println(timeElement.currentTime.String())
}
