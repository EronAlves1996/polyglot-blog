package main

import (
	"log"
	"net/http"
)

var hostPort = ":4000"

type HandlerMatcher struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func (hm *HandlerMatcher) register(method string,
	handler func(http.ResponseWriter, *http.Request)) {
	hm.handlers[method] = handler
}

func (hm *HandlerMatcher) tryMatch(w http.ResponseWriter,
	r *http.Request) {
	handler := hm.handlers[r.Method]
	if handler == nil {
		http.Error(w, "Method not Allowed", 405)
		return
	}
	handler(w, r)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Error(w, "Not Found!", 404)
		return
	}

	hm := HandlerMatcher{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	hm.register(http.MethodGet, func(wr http.ResponseWriter,
		r *http.Request) {
		w.Write([]byte("Placeholder for home page"))
	})

	hm.tryMatch(w, r)
}

func search(w http.ResponseWriter, r *http.Request) {
	hm := HandlerMatcher{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placheolder for search page"))
	})

	hm.tryMatch(w, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/search", search)

	err := http.ListenAndServe(hostPort, mux)
	log.Fatal(err)
}
