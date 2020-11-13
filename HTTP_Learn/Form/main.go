package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	r.ParseForm()
	// 获取请求参数
	fmt.Fprintln(w, "请求参数:", r.Form)
	fmt.Fprintln(w, "请求参数", r.PostForm)
	fmt.Fprintln(w, "请求参数:", r.Form.Get("key"))
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8083", nil)
}
