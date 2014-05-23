package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

/*
   markov chain text generation
   http;//golang/org/doc/codewalk/markov/
*/

func Bot() io.ReadWriteCloser {
	r, out := io.Pipe() // for outgoing data
	return bot{r, out}
}

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	go b.speak()
	return len(buf), nil
}

func (b bot) speak() {
	time.Sleep(time.Second)
	msg := "test" // chain.Generate(10) // at most 10 words
	b.out.Write([]byte(msg))
}

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Waiting for a partner...")
	select {
	case partner <- c:
		// handled by the second goroutine
		fmt.Fprintln(c, "Delegating you to a partner...")
	case p := <-partner:
		chat(p, c)
	case <-time.After(10 * time.Second):
		chat(Bot(), c)
	}
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found a partner! Say hi.")
	fmt.Fprintln(b, "Found a partner! Say hi.")
	errc := make(chan error, 1)
	go cp(a, b, errc)
	go cp(b, a, errc)

	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	a.Close()
	b.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go match(c)
	}
}
