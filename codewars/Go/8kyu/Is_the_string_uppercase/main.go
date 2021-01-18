package main

import (
	"fmt"
	"strconv"
)

type MyString string

func main() {
		MyString("StartTrek").IsUpperCase()
}

func (s MyString) IsUpperCase() bool {
	var o bool = true
	var s string
	for i, c := range s {
		v,_ := fmt.Println(i, c)
		v,_ = strconv.Atoi(string(v))
		if v >= 65 && v <= 90 {
			s:=fmt.Sprint(o)
		}
		return s
	}
}