package main

import (
	"net/http"
)

func main() {
	go http2()
	http1()

}
func http1() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("/Users/alphayan/go/src/github.com/gopl-zh/gopl-zh.github.com/_book")))
	http.ListenAndServe(":9001", mux)
}
func http2() {
	http.Handle("/", http.FileServer(http.Dir("/Users/alphayan/go/src/github.com/chai2010/advanced-go-programming-book/_book")))
	http.ListenAndServe(":9000", nil)
}
