package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type User struct {
	Name string
	Ear  chan string
}

func (self User) Listen() {
	for {
		select {
		case msg := <-self.Ear:
			self.Say(fmt.Sprintf("Hey, I've been mentioned in '%s'!", msg))
		}
	}
}

func (self User) Say(msg string) {
	log.Printf("%s: %s\n", self.Name, msg)
}

func mkUser(name string) User {
	return User{
		Name: name,
		Ear:  make(chan string),
	}
}

func stripAts(strs []string) (result []string) {
	result = make([]string, len(strs))
	for i, str := range strs {
		result[i] = str[1:]
	}

	return
}

func mentions(msg string) []string {
	re := regexp.MustCompile("@(\\w+)")
	return stripAts(re.FindAllString(msg, 140))
}

func processMentions(msg string) {
	for _, mention := range mentions(msg) {
		if users[mention].Ear != nil {
			users[mention].Ear <- msg
		}
	}
}

var users map[string]User

func main() {
	users = map[string]User{
		"Mark": mkUser("Mark"),
		"Tom":  mkUser("Tom"),
		"Bob":  mkUser("Bob"),
	}

	for _, user := range users {
		go user.Listen()
	}

	you := mkUser("You")
	fmt.Println("Say something, mention Mark, Bob or Tom with '@'. Say 'quit' to exit.")

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text == "quit" {
			break
		}

		you.Say(text)
		processMentions(text)
	}

	for _, user := range users {
		close(user.Ear)
	}
}
