package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "mux", r.URL.Path)
}

func main() {
 	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
 	http.ListenAndServe(":8889", mux)
}
