package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "发送：", r.URL)
	fmt.Fprintln(w, "发送", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头所有信息", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头所有信息", r.Header.Get("Cookie"))
	fmt.Fprintln(w, "请求体", r.ContentLength)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体", string(body))

}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8083", nil)
}