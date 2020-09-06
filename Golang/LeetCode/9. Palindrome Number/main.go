package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	length := 0
	for temp := x; temp > 0; temp = temp / 10 {
		length++
	}
	temp := length / 2
	reverse := 0
	for i := 0; i < temp; i++ {
		k := x % 10
		x = x / 10
		reverse = reverse*10 + k
	}
	if length%2 == 1 {
		x = x / 10
	}
	if x == reverse {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Print(isPalindrome(12321))
}
