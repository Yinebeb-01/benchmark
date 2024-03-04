package main

import (
	"fmt"
	"regexp"
	"strings"
)

func regEx(pass string) bool {
	re, err := regexp.Compile("[0123456789]")
	if err != nil {
		fmt.Println("error occurred ", err)
		return false
	}
	if !re.Match([]byte(pass)) {
		return false
	}
	return true
}

func strContains(password string) bool {
	return strings.ContainsAny(password, "0123456789")
}
