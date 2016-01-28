package main

import (
    "net/url"
    "fmt"
)

type Validator struct {
    config *Configuration
}

// Checks if given request image host belongs to one in the
// white list.
func (v *Validator) CheckHostInWhiteList(requestUrl string) error {
    urlParsed, err := url.Parse(requestUrl)

    if err != nil {
        return err
    }

    var hostFound bool

    for _, host := range v.config.DomainWhiteList {
        if host == urlParsed.Host {
            hostFound = true
        }
    }

    if hostFound {
        return nil
    }

    return error(fmt.Errorf("Host %s not allowed", urlParsed.Host))
}

// Validates if new request size is valid or not
func (v *Validator) CheckRequestNewSize(s *Size) error {
    if s.Height <= 0 || s.Width <= 0 {
        return error(fmt.Errorf("Width or height should be bigger than 0"))
    }

    if  s.Height >= v.config.Size.Height {
        return error(fmt.Errorf("Height cannot be higher than %d", v.config.Size.Height))
    }

    if s.Width >= v.config.Size.Width {
        return error(fmt.Errorf("Width cannot be higher than %d", v.config.Size.Width))
    }

    return nil
}