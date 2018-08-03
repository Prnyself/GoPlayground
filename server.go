package main

import (
	"net/http"
	"fmt"
	"path"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:3009",
		Handler: nil,
	}
	http.HandleFunc("/", newHandle)
	server.ListenAndServe()
}

func newHandle(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintln(writer, path.Base(request.URL.Path))
}
