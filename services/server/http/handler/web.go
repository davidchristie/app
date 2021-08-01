package handler

import (
	"net/http"
	"os"
)

func Web(dir string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(dir + r.URL.Path); os.IsNotExist(err) {
			http.StripPrefix(r.URL.Path, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	}
}
