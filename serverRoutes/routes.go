package serverRoutes

import (
	"net/http"
	"fmt"
)

var MyRouter *http.ServeMux

func init() {
	MyRouter = http.NewServeMux()
	MyRouter.HandleFunc("/hello", sayHello)
	MyRouter.HandleFunc("/test", anotherTest)
}
func anotherTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world again")
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
