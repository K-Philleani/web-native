package main

import (
	"fmt"
	"net/http"
)

// 处理器
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8087", nil)
}
