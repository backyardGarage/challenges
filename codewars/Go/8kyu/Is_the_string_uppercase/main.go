package main

import (
	"fmt"
)

type MyString string

func main() {
		MyString("StartTrek").IsUpperCase()
}

func (s MyString) IsUpperCase() bool {
	var t bool = true
	for _, c := range s {
		if c >= 97 && c <= 122 {
			t = false
			return t
		}
	}
	return t
}