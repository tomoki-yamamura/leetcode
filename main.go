package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/ziip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
	file, err := os.Create("ascii_sample.zip")
	if err != nil {
		panic(err)
	}
	io.Copy(w, file)
}