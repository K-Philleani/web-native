package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8085", nil)
}