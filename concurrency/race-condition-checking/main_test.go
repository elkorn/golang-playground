package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestWithRace(t *testing.T) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	run(withRaceCondition)
	fmt.Println("Ran without synchronization.")
}

func TestSynchronizedWithMutex(t *testing.T) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	run(synchronizedByChannels)
	fmt.Println("Ran synchronized by channels.")
}

func TestSynchronizedWithChannels(t *testing.T) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	run(synchronizedByChannels)
	fmt.Println("Ran synchronized by channels.")
}
