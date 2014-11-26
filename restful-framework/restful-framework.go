package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// type url.Values map[string][]string - how it is defined in the net/url package.
type Resource interface {
	Get(values url.Values) (int, interface{}) // Return a status code and the data in any format.
	// interface{} means any type.
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

func notSupported() (int, interface{}) {
	return 405, ""
}

type (
	API                struct{}
	GetNotSupported    struct{}
	PostNotSupported   struct{}
	PutNotSupported    struct{}
	DeleteNotSupported struct{}
	GetOnlyResource    struct {
		PostNotSupported
		PutNotSupported
		DeleteNotSupported
	}
)

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return notSupported()
}

func (PostNotSupported) Post(values url.Values) (int, interface{}) {
	return notSupported()
}

func (PutNotSupported) Put(values url.Values) (int, interface{}) {
	return notSupported()
}

func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
	return notSupported()
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		method := req.Method
		values := req.Form

		var data interface{}
		var code int

		switch method {
		case "GET":
			code, data = resource.Get(values)
		case "POST":
			code, data = resource.Post(values)
		case "PUT":
			code, data = resource.Put(values)
		case "DELETE":
			code, data = resource.Delete(values)
		default:
			api.Abort(rw, 405)
			return
		}

		content, err := json.Marshal(data)
		if nil != err {
			log.Println(err)
			api.Abort(rw, 500)
			return
		}

		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (api *API) AddResource(resource Resource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (api *API) Abort(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
}

type Home struct {
	GetOnlyResource
}

type Test struct {
	a int
	b bool
}

func (home Home) Get(values url.Values) (int, interface{}) {
	fmt.Println("Home:Get")
	return 200, Test{a: 13, b: true}
}

func main() {
	api := API{}
	api.AddResource(Home{}, "/home")
	api.Start(3456)
}
