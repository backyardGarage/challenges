// Array training
// source: https://ehpub.co.kr/go-%EC%96%B8%EC%96%B4-30-%EB%B0%B0%EC%97%B4%EC%9D%98-%EC%9A%94%EC%86%8C%EB%A5%BC-%EC%88%9C%EC%B0%A8%EC%A0%81%EC%9C%BC%EB%A1%9C-%EB%B0%A9%EB%AC%B8%ED%95%98%EA%B8%B0/

package main

import "fmt"

func main() {
	arr := [5]int{12, 34, 23, 56, 34}

	for i := range arr {
		fmt.Println(i)
	}
}
