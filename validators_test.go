package main

import (
    "testing"
)

func TestCheckHostInWhiteListWithEmptyConfiguration(t *testing.T) {
    config := new(Configuration)
    validator := Validator{config}

    if err := validator.CheckHostInWhiteList("doesnt exists"); err == nil {
        t.Errorf("Missing error returning!")
    }
}

func TestCheckHostInWhiteListWithSomeHostsInWhieList(t *testing.T) {
    config := new(Configuration)
    config.HostWhiteList = []string{"http://www.google.com", "two hosts"}
    validator := Validator{config}

    // Check for one that doesn't exists
    err := validator.CheckHostInWhiteList("http://www.sergiosola.com/")
    if err == nil {
        t.Errorf("Should return an error!!!")
    }
}

func TestCheckHostInWhiteListWithValidHost(t *testing.T) {
    config := new(Configuration)
    config.HostWhiteList = []string{"one host", "http://sergiosola.com"}
    validator := Validator{config}

    err := validator.CheckHostInWhiteList("https://sergiosola.com/images?withParams=dsada")
    if err == nil {
        t.Errorf("Should not to return an error!!!")
    }
}