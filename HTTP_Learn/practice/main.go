package main

import (
	"fmt"
	"net/http"
)

// 处理器函数方式
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "处理器函数", r.URL.Path)
//}
//
//func main() {
//	http.HandleFunc("/handler", handler)
//	http.ListenAndServe(":9091", nil)
//}

type MyHandler struct {}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "处理器", r.URL.Path)
}

func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr: ":9091",
		Handler: &myHandler,
	}
	server.ListenAndServe()
}
