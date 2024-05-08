package main

import (
	"html/template"
	"log"
	"net/http"
)

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

func newMatcher() HandlerMatcher {
	return HandlerMatcher{
		handlers: map[string]func(http.ResponseWriter, *http.Request){},
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Error(w, "Not Found!", 404)
		return
	}

	hm := newMatcher()

	hm.register(http.MethodGet, func(wr http.ResponseWriter,
		r *http.Request) {
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			log.Print(err)
			http.Error(w, "Internal server error", 500)
			return
		}
		ts.Execute(w, nil)
	})

	hm.tryMatch(w, r)
}

func search(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placheolder for search page"))
	})

	hm.tryMatch(w, r)
}

func viewPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/post" {
		http.Error(w, "Not Found!", 404)
		return
	}

	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placholder for blog post page"))
	})

	// create a post
	hm.register(http.MethodPost, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Location", "/fake/location")
		w.Write([]byte("Placeholder for creating blog post page"))
	})

	hm.register(http.MethodPut, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placeholder for editing a blog post"))
	})

	hm.register(http.MethodDelete, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placeholder for deleting a blog post"))
	})

	hm.tryMatch(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placeholder for login a user"))
	})

	hm.tryMatch(w, r)
}

func register(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodPost, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Location", "/fake/location")
		w.Write([]byte("Placeholder for registering a new user"))
	})

	hm.tryMatch(w, r)
}
