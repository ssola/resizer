package main

import (
    "fmt"
    "net/http"
    "io"
    "github.com/nfnt/resize"
    "image"
    "image/jpeg"
    "image/png"
    "strconv"
)

// Return a given error in JSON format to the ResponseWriter
func format_error(err error, w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, fmt.Sprintf("{ \"error\": \"%s\"}", err))
    return
}

// Parse a given string into a uint value
func parseInteger(value string) (uint, error) {
    integer, err := strconv.Atoi(value)
    return uint(integer), err
}

// Resizing endpoint.
func resizing(w http.ResponseWriter, r *http.Request) {
    var newWidth, newHeight uint

    // Get parameters
    imageUrl := r.FormValue("image")
    newWidth, _ = parseInteger(r.FormValue("width"))
    newHeight, _ = parseInteger(r.FormValue("height"))

    // Download the image
    imageBuffer, err := http.Get(imageUrl)
    if err != nil {
        format_error(err, w)
    }

    finalImage, _, _ := image.Decode(imageBuffer.Body)

    r.Body.Close()

    imageResized := resize.Resize(newWidth, newHeight, finalImage, resize.Lanczos3)

    if imageBuffer.Header.Get("Content-Type") == "image/png" {
        png.Encode(w, imageResized)
    }

    if imageBuffer.Header.Get("Content-Type") == "image/jpg" {
        jpeg.Encode(w, imageResized, nil)
    }

    if imageBuffer.Header.Get("Content-Type") == "binary/octet-stream" {
        jpeg.Encode(w, imageResized, nil)
    }
}

func main() {
    http.HandleFunc("/resize", resizing)
    http.ListenAndServe(":8080", nil)
}
