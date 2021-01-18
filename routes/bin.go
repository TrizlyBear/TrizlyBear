package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Bin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./public/bin/static/index.html")
	case "POST":
		fmt.Println("IDK")
		keys := []string{"test123"}
		var p binfile
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(p)
		if _, err := os.Stat("./public/bin/static/" + p.Name + "." + p.Code); !os.IsNotExist(err) {
			http.Error(w, "File already exists", http.StatusBadRequest)
		} else if p.Value == "" {
			http.Error(w, "Please provide code", http.StatusBadRequest)
		} else if !stringInSlice(p.Key, keys) {
			http.Error(w, "Provide a (working) key", http.StatusBadRequest)
		} else {
			err = ioutil.WriteFile("./public/bin/static/"+p.Name+"."+p.Code, []byte(p.Value), 0644)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, p.Name+"."+p.Code)
		}
	}

}

type binfile struct {
	Code  string
	Key   string
	Name  string
	Value string
}

type FileSystem struct {
	FS http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.FS.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := "./." + strings.TrimSuffix(path, "/") + "/index.html"
		fmt.Println(index)
		if _, err := fs.FS.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
