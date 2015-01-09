package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
