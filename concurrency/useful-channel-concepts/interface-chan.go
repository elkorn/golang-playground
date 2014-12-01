package main

import (
	"fmt"
)

type Messenger interface {
	Relay() string
}

type Message struct {
	Content string
}

func (m Message) Relay() string {
	return m.Content
}

func alertMessage(v chan Messenger, i int) {
	m := new(Message)
	m.Content = fmt.Sprintf("Done with %d", i)
	v <- m
}

func showInterfaceChannel() {
	msg := make(chan Messenger)

	for i := 0; i < 10; i++ {
		go alertMessage(msg, i)
	}

	for i := 0; i < 10; i++ {
		select {
		case message := <-msg:
			fmt.Println(message.Relay())
		}
	}
}
