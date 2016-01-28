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
    config.DomainWhiteList = []string{"one host", "two hosts"}
    validator := Validator{config}

    // Check for one that doesn't exists
    err := validator.CheckHostInWhiteList("one host2")
    if err.Error() != "Host  not allowed" {
        t.Errorf("Should return an error!!!")
    }
}