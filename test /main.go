package main1

import (
	"log"
	"net/http"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello world!"))
}

func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
