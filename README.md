# httprouter
Simple http router, example:

```
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
	server := &Server{Router: httprouter.NewRouter()}
	server.SetIfNotMatch(notMatch)
	server.Get("/echo", echo)
	return server
}

func main() {
	server := NewServer()
	fmt.Println(http.ListenAndServe(":8080", server))
}
```
