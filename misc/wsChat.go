package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket" // Use `go get` to download this package.
)

// This is not finished. more details are available in the chatroulette presentation.

var rootTemplate = template.Must(template.New("Root").Parse(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8" />
		</head>
		<script>
			var websocket = new WebSocket("ws://{{.}}/socket");
			websocket.onmessage = function() {
				console.log(arguments);
			};

			websocket.onclose = function() {
				console.log(arguments);
			};
		</script>
	</html>
	`))

var partner = make(chan io.ReadWriteCloser)

type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, "localhost:4000")
}

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Waiting for a partner...")
	select {
	case partner <- c:
		// handled by the second goroutine
		fmt.Fprintln(c, "Delegating you to a partner...")
	case p := <-partner:
		chat(p, c)
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
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
