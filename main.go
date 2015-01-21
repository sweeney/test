package main

import (
	"fmt"
	"net/http"
	//"time"
)

func main() {
	serve()
	return
}

func serve() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", hw)
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func hw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
