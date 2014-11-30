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

func TestSynchronizedWithChannels(t *testing.T) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	run(synchronizedByChannels)
	fmt.Println("Ran synchronized by channels.")
}
