package httprouter

import (
	"fmt"
	"net/http"
	"testing"
)

func notMatch(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("not match any handler")
}

func echo(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("match echo handler")
}

func TestRouter(t *testing.T) {
	router := newRouter()
	router.SetIfNotMatch(notMatch)
	router.Get("/echo", echo)
	router.Match("GET", "/echo")(nil, nil)
	router.Match("POST", "/echo")(nil, nil)
	router.Match("GET", "/not_exist")(nil, nil)
}
