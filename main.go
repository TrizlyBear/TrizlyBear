package main

import (
	"fmt"
	"github.com/TrizlyBear/TrizlyBear/routes"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	PORT = "localhost:6969"
)

func main() {
	fmt.Println("Running on port %v", PORT)
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	//fs := http.FileServer(routes.FileSystem{http.Dir("./public/bin/static")})
	//r.PathPrefix("/bin/").Handler(http.StripPrefix("/bin/", fs))
	r.HandleFunc("/bin", routes.Bin)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Golang 😎", r.URL.Path[1:])
}
