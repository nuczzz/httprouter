package main

import (
	"fmt"
	"net/http"

	"github.com/nuczzz/httprouter"
)

func notMatch(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "not match any handler.")
}

func echo(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "match echo handler")
}

type Server struct {
	httprouter.Router
}

func NewServer() *Server {
	router := httprouter.NewRouter()
	router.SetIfNotMatch(notMatch)
	router.Get("/echo", echo)
	return &Server{Router: router}
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.Match(req.Method, req.URL.Path)(resp, req)
}

func main() {
	server := NewServer()
	fmt.Println(http.ListenAndServe(":8080", server))
}
