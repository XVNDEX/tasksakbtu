package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health", foo)
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
}