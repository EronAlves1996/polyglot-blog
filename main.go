package main

import (
	"log"
	"net/http"
)

var hostPort = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Error(w, "Not Found!", 404)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not Allowed", 405)
		return
	}

	w.Write([]byte("Placeholder for home page"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(hostPort, mux)
	log.Fatal(err)
}
