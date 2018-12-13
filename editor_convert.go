package main

import "net/http"

type myHandler struct {
}

func (h myHandler) isChrome(r *http.Request) bool {
	return true
}

func (h myHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if h.isChrome(r) {
		w.Write([]byte("This is Chrome instance"))
	}
	w.Write([]byte("This is indexHandler"))
}

func main() {

}
