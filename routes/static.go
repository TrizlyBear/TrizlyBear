package routes

import (
	"net/http"
)

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
		return nil, nil
	}

	return f, nil
}
