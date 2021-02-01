/*
The vowel substrings in the word codewarriors are o,e,a,io.
The longest of these has a length of 2.
Given a lowercase string that has alphabetic characters only (both vowels and consonants) and
no spaces, return the length of the longest vowel substring.
Vowels are any of aeiou.
*/

// 특정 문자열이 최대 몇개 까지 연속되어 포함되어있는지를 확인할 수 있어야함
// 나의 로직
// 일단 받은 모든 문자열에서 모음에 해당하는 문자열을 지정함
// 해당 문자열의 index를 더했을 때, 더한 값의 문자열 또한 모음일 경우,
// 연속모음 카운터 변수를 증가시킴.
// 만약 모음이 아닐 경우, 비교대상 인덱스를 증가시킴

// 해야할 일 -> 입력받은 문자열을 슬라이스에 집어넣기
package main

import (
	"fmt"
	"strings"
)

func Solve(s string) int {
	s = strings.ToLower(s)
	fmt.Printf("%s", s)

	return 0
}

func main() {
	Solve("codewarriors")
}
