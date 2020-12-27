package main

func CorrectTail(body string, tail rune) bool {
	t := len(body) - 1
	if rune(body[t]) == tail {
		return true
	}
	return false
}
