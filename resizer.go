package main

import (
    "fmt"
    "net/http"
    "github.com/nfnt/resize"
    "image"
    "image/jpeg"
    "image/png"
    "strconv"
)

type Size struct {
    Width uint
    Height uint
}

func validateSize(s *Size) error {
    if s.Height >= 1000 {
        return error(fmt.Errorf("Height cannot be higher than 1000"))
    }

    if s.Width >= 1000 {
        return error(fmt.Errorf("Width cannot be higher than 1000"))
    }

    return nil
}

// Return a given error in JSON format to the ResponseWriter
func formatError(err error, w http.ResponseWriter)  {
    http.Error(w, fmt.Sprintf("{ \"error\": \"%s\"}", err), 400)
}

// Parse a given string into a uint value
func parseInteger(value string) (uint, error) {
    integer, err := strconv.Atoi(value)
    return uint(integer), err
}

// Resizing endpoint.
func resizing(w http.ResponseWriter, r *http.Request) {
    size := Size{}

    // Get parameters
    imageUrl := r.FormValue("image")
    size.Width, _ = parseInteger(r.FormValue("width"))
    size.Height, _ = parseInteger(r.FormValue("height"))

    if err := validateSize(&size); err != nil {
        formatError(err, w)
        return
    }

    // Download the image
    imageBuffer, err := http.Get(imageUrl)
    if err != nil {
        formatError(err, w)
        return
    }

    finalImage, _, _ := image.Decode(imageBuffer.Body)

    r.Body.Close()

    imageResized := resize.Resize(size.Width, size.Height, finalImage, resize.Lanczos3)

    if imageBuffer.Header.Get("Content-Type") == "image/png" {
        png.Encode(w, imageResized)
    }

    if imageBuffer.Header.Get("Content-Type") == "image/jpeg" {
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
