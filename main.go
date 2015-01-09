package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/_hw", hw)
	http.HandleFunc("/_ping", ping)
	http.ListenAndServe(":8080", nil)
}

func hw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helloworld")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
