package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	sum := ""
	index1 := len1 - 1
	index2 := len2 - 1
	carry := 0
	for index1 >= 0 && index2 >= 0 {
		n1 := getInt(num1[index1])
		n2 := getInt(num2[index2])
		sum = strconv.Itoa((n1+n2+carry)%10) + sum
		carry = (n1 + n2 + carry) / 10
		index1--
		index2--
	}
	if index1 < 0 {
		for index2 >= 0 {
			n2 := getInt(num2[index2])
			sum = strconv.Itoa((n2+carry)%10) + sum
			carry = (n2 + carry) / 10
			index2--
		}
	} else if index2 < 0 {
		for index1 >= 0 {
			n1 := getInt(num1[index1])
			sum = strconv.Itoa((n1+carry)%10) + sum
			carry = (n1 + carry) / 10
			index1--
		}
	}

	if index1 < 0 && index2 < 0 {
		if carry == 0 {
			return sum
		} else {
			return strconv.Itoa(carry) + sum
		}
	}
	return sum
}

func getInt(num uint8) int {
	return int(num - 48)
}

func main() {
	num1 := "919"
	num2 := "99"
	sum := addStrings(num1, num2)
	fmt.Println(sum)
}
