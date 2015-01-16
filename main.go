package main

import (
	"fmt"
	"net/http"
	//"time"
)

func main() {
	//go serve()
	serve()
	//time.Sleep(30 * time.Second)
	return
}

func serve() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
