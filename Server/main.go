package main

import (
	"fmt"
	"net/http"
	"time"
)

type ServerHandler struct {

}
func (s *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is a server", r.URL.Path)
}


func main() {
	serverHandler := ServerHandler{}

	server := http.Server{
		Addr:              ":8888",
		Handler:           &serverHandler,
		TLSConfig:         nil,
		ReadTimeout:       2 * time.Second,
	}

	server.ListenAndServe()
}
