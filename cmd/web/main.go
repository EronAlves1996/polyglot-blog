package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Port where server gonna listen. Default: ':4000'")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/search", search)
	mux.HandleFunc("/post/", viewPost)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
