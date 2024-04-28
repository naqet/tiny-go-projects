package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello"))
	})


	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe("127.0.0.1:8000", mux)
}
