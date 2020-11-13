package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL user:", r.FormValue("user"))
	fmt.Fprintln(w, "URL user:", r.PostFormValue("user"))
}


func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8083", nil)
}
