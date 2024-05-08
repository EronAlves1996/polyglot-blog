package main

import (
	"log"
	"net/http"
)

var hostPort = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/search", search)
	mux.HandleFunc("/post/", viewPost)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)

	err := http.ListenAndServe(hostPort, mux)
	log.Fatal(err)
}
