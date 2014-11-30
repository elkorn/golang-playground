package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /register")
	params := mux.Vars(r)
	name := params["name"]

}

func main() {

}
