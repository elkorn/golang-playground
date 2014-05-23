package main

//http://tour.golang.org/#54
import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func main() {
	//var h Hello
	http.Handle("/string", String("Handling string."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
    
}
