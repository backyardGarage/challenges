package main

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {

	var arrUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var space string = " "
	var a = strings.Contains(string(s), arrUpper)
	var b = strings.Contains(string(s), space)

	return b && a
}

func main() {
	s := "Hello"
	t := MyString.IsUpperCase(MyString(s))
	fmt.Println(t)
}
