package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is HandleFunc", r.URL.Path)
}

func main() {
	http.HandleFunc("/handleFunc", handler)
	http.ListenAndServe(":8091", nil)
}