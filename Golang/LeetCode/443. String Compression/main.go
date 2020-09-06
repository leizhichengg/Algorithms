package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	res := 0
	i := 0
	for i < len(chars) {
		current := chars[i]
		count := 0
		for i < len(chars) && chars[i] == current {
			i++
			count++
		}
		chars[res] = current
		res++
		if count > 1 {
			copy(chars[res:], strconv.Itoa(count))
			res += len(strconv.Itoa(count))
		}
	}
	return res
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	fmt.Println(chars)
}
