package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	PORT = ":6969"
)

func main() {
	fmt.Println("Running on port %v", PORT)

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Golang 😎", r.URL.Path[1:])
}
