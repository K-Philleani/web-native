package main

import (
	"fmt"
	"net/http"
)

type Handle struct {}

func (h *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is a handle", r.URL.Path)
}

func main() {
	handle := Handle{}
	http.Handle("/handle", &handle)
	http.ListenAndServe(":8092", nil)
}
