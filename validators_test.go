package main

import "testing"

func TestCheckHostInWhiteListWithEmptyConfiguration(t *testing.T) {
    config := new(Configuration)
    validator := Validator{config}

    if err := validator.CheckHostInWhiteList("doesnt exists"); err == nil {
        t.Errorf("Missing error returning!")
    }
}