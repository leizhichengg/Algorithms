package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	temp := map[[26]byte][]string{}
	for _, str := range strs {
		strByte := getCharByte(str)
		if _, ok := temp[strByte]; ok {
			temp[strByte] = append(temp[strByte], str)
		} else {
			temp[strByte] = []string{str}
		}
	}
	ans := [][]string{}
	for _, v := range temp {
		ans = append(ans, v)
	}
	return ans
}

func getCharByte(str string) [26]byte {
	res := [26]byte{}
	for _, s := range str {
		res[s-'a']++
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}
