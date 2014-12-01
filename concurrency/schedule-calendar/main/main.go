package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/elkorn/golang-playground/concurrency/schedule-calendar/domain"
	"github.com/elkorn/golang-playground/concurrency/schedule-calendar/presentation"
	"github.com/gorilla/mux"
)

var mutex sync.Mutex

func initUser(user *domain.User) {
	for i := 9; i < 18; i++ {
		user.Times[i] = true
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /register")
	params := mux.Vars(r)
	name := params["name"]
	var page *presentation.Page
	if _, ok := domain.Users[name]; ok {
		page = presentation.MkPage("User already exists", fmt.Sprintf("User %s already exists", name), nil)

	} else {
		newUser := domain.User{
			Name: name,
		}

		initUser(&newUser)
		domain.Users[name] = newUser

		page = presentation.MkPage("User created.", fmt.Sprintf("You have created user %s.", name), nil)
	}

	page.Render("presentation/generic.txt", w)
}

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /users")
	page := presentation.MkPage("View users", "", domain.Users)
	page.Render("presentation/generic.txt", w)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /schedule")
	params := mux.Vars(r)
	name := params["name"]
	time := params["time"]

	timeVal, _ := strconv.ParseInt(time, 10, 0)
	timeVal = int(timeVal)

	createURL := "/register" + name
	if _, ok := domain.Users[name]; ok {
		mutex.Lock()
		domain.Users[name].Times[intTimeVal] = false
		mutex.Unlock()
		// Ok, I have enough. This is the only thing in this crap example that 
		// has anything to do with concurrency.
	} else {
	}
}

func main() {

}
