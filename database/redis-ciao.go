package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alphazero/Go-Redis"
)

func main() {
	testSpec := redis.DefaultSpec().Db(13)
	client, e := redis.NewSynchClientWithSpec(testSpec)

	if nil != e {
		log.Println("failed to create the client", e)
		return
	}

	key := "examples/hello/user.name"
	value, e := client.Get(key)
	if nil != e {
		log.Println("error on Get", e)
		return
	}

	if nil == value {
		fmt.Printf("\nhello! Gimme your name.")
		reader := bufio.NewReader(os.Stdin)
		user, _ := reader.ReadString(byte('\n'))

		if len(user) > 1 {
			user = user[0 : len(user)-1]
			value = []byte(user)
			client.Set(key, value)
		} else {
			fmt.Printf("vafanculo!\n")
			return
		}
	}

	fmt.Printf("Hey, ciao %s!\n", fmt.Sprintf("%s", value))
	select {
	case <-time.After(60 * time.Second):
	}
}
