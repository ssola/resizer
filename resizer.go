package main

import (
    "fmt"
    "net/http"
    "github.com/nfnt/resize"
    "image"
    "image/jpeg"
    "image/png"
    "strconv"
    "github.com/spf13/viper"
)

type Configuration struct {
    Port uint
    HostWhiteList []string
    Size Size
}

type Size struct {
    Width uint
    Height uint
}

var config *Configuration

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
    size := new(Size)

    // Get parameters
    imageUrl := r.FormValue("image")
    size.Width, _ = parseInteger(r.FormValue("width"))
    size.Height, _ = parseInteger(r.FormValue("height"))

    validator := Validator{config}

    if err := validator.CheckHostInWhiteList(imageUrl); err != nil {
        formatError(err, w)
        return
    }

    if err := validator.CheckRequestNewSize(size); err != nil {
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
    // Load configuration
    viper.SetConfigName("config")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        panic(fmt.Errorf("Fatal error loading configuration file: %s", err))
    }

    // Marshal the configuration into our Struct
    viper.Unmarshal(&config)

    http.HandleFunc("/resize", resizing)
    http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
