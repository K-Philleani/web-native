package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is HandleFunc", r.URL.Path)
}

func main() {
	http.HandleFunc("/start", Handler)
	http.ListenAndServe(":9090", nil)
}
