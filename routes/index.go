package routes

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/static/comming-soon/index.html")
}
