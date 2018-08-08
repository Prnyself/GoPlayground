package main

import (
	"net/http"
	"fmt"
	"path"
	"awesomeProject/serverRoutes"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:3009",
		Handler: serverRoutes.MyRouter,
	}
	//fmt.Println(server)
	//http.HandleFunc("/aaa", newHandle)
	//http.HandleFunc("/hi", sayHi)
	server.ListenAndServe()
	//http.ListenAndServe(":3009", nil) // 注意这里addr只输端口号的话不能漏了冒号
}

func newHandle(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintln(writer, path.Base(request.URL.Path))
}


func sayHi(writer http.ResponseWriter, request *http.Request){
	params := request.URL.Query()
	//name := params.Get("name")
	name := params["name"]
	SayHi(1)
	fmt.Fprintln(writer, name)
}

func SayHi(a int) int {
	return a
}

