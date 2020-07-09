package main

import "fmt"

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	length := len(s)
	mid := length / 2
	for i := 0; i < mid; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}
