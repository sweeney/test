package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hw)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}

func hw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helloworld")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
