package main

import "time"

func now() int64 {
	return time.Now().UnixNano()
}
