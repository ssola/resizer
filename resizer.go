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
    Size Size
}

type Size struct {
    Width uint
    Height uint
}

var config *Configuration

// Including basic validation to prevent creating big images
func validateSize(c *Configuration, s *Size) error {
    if  s.Height >= c.Size.Height {
        return error(fmt.Errorf("Height cannot be higher than %d", c.Size.Height))
    }

    if s.Width >= c.Size.Width {
        return error(fmt.Errorf("Width cannot be higher than %d", c.Size.Width))
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
    size := new(Size)

    // Marshal the configuration into our Struct
    viper.Unmarshal(&config)

    // Get parameters
    imageUrl := r.FormValue("image")
    size.Width, _ = parseInteger(r.FormValue("width"))
    size.Height, _ = parseInteger(r.FormValue("height"))

    if err := validateSize(config, size); err != nil {
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

    http.HandleFunc("/resize", resizing)
    http.ListenAndServe(":8080", nil)
}
