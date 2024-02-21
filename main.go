package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)

	http.ListenAndServe(":3333", mux)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
