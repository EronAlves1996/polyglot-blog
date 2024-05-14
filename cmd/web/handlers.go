package main

import (
	"bytes"
	"fmt"
	"html/template"
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

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Error(w, "Not Found!", 404)
		return
	}

	hm := newMatcher()

	hm.register(http.MethodGet, func(wr http.ResponseWriter,
		r *http.Request) {
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			a.serverError(err, wr)
			return
		}
		ts.Execute(w, nil)
	})

	hm.tryMatch(w, r)
}

func (a *app) search(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("s")
		w.Write([]byte(fmt.Sprintf("Placheolder for search page, searching %s", query)))
	})

	hm.tryMatch(w, r)
}

func (a *app) viewPost(w http.ResponseWriter, r *http.Request) {
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

type NewPostTemplateData struct {
	Title string
}

func (a *app) newPost(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/html/base.html",
			"./ui/html/new-post.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			a.serverError(err, w)
			return
		}

		bf := new(bytes.Buffer)

		td := NewPostTemplateData{
			Title: "Novo Post",
		}

		if err := ts.ExecuteTemplate(bf, "base", td); err != nil {
			a.serverError(err, w)
			return
		}
		w.WriteHeader(http.StatusOK)
		bf.WriteTo(w)
	})

	hm.tryMatch(w, r)
}

func (a *app) login(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Placeholder for login a user"))
	})

	hm.tryMatch(w, r)
}

func (a *app) register(w http.ResponseWriter, r *http.Request) {
	hm := newMatcher()

	hm.register(http.MethodPost, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Location", "/fake/location")
		w.Write([]byte("Placeholder for registering a new user"))
	})

	hm.tryMatch(w, r)
}
