package main

import "fmt"

func hammingWeight(num uint32) int {
	var ans uint32
	for num != 0 {
		ans += num & 1
		num = num >> 1
	}
	return int(ans)
}

func main() {
	fmt.Println(hammingWeight(9))
}
