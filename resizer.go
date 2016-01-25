package main

import (
	"fmt"
	"net/http"
	"io"
)

func format_error(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf("{ \"error\": \"%s\"}", err))
	return
}

func resize(w http.ResponseWriter, r *http.Request) {
	originalImage := r.FormValue("image")
	// newWidth := r.FormValue("width")
	// newHeight := r.FormValue("height")

	// Download the image
	image, err := http.Get(originalImage)
	if err != nil {
		format_error(err, w)
	}

	// If we have the image then let's try to return it back.
	r.Header.Set("Content-Length", fmt.Sprint(image.ContentLength))
	r.Header.Set("Content-Type", r.Header.Get("Content-Type"))

	// We just copy the content from the image to the writer.
	if _, err = io.Copy(w, image.Body); err != nil {
		format_error(err, w)
	}

	r.Body.Close()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", resize)
	http.ListenAndServe(":8080", mux)
}
