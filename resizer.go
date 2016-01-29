package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

type Configuration struct {
	Port          uint
	HostWhiteList []string
	Size          Size
}

type Size struct {
	Width  uint
	Height uint
}

var config *Configuration

// Return a given error in JSON format to the ResponseWriter
func formatError(err error, w http.ResponseWriter) {
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

	contentType := imageBuffer.Header.Get("Content-Type")
	switch contentType {
	case "image/png":
		png.Encode(w, imageResized)
		log.Printf("Successfully handled content type '%s'\n", contentType)
	case "image/jpeg":
		jpeg.Encode(w, imageResized, nil)
		log.Printf("Successfully handled content type '%s'\n", contentType)
	case "binary/octet-stream":
		jpeg.Encode(w, imageResized, nil)
		log.Printf("Successfully handled content type '%s'\n", contentType)
	default:
		log.Printf("Cannot handle content type '%s'\n", contentType)
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
