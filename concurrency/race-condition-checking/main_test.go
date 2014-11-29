package main

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestWithRace(t *testing.T) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	run(withRaceCondition)
}
