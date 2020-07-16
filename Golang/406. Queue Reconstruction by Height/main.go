package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	for index, p := range people {
		to := p[1]
		copy(people[to+1:index+1], people[to:index])
		people[to] = p
	}
	return people
}

func main() {
	people := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	ans := reconstructQueue(people)
	fmt.Println(ans)
}
