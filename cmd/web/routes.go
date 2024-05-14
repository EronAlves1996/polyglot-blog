package main

import "net/http"

func (a *app) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/search", a.search)
	mux.HandleFunc("/post/", a.viewPost)
	mux.HandleFunc("/posts/new", a.newPost)
	mux.HandleFunc("/login", a.login)
	mux.HandleFunc("/register", a.register)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	return mux
}
