package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return construct(1, n)
}

func construct(start, end int) []*TreeNode {
	ans := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := construct(start, i-1)
		right := construct(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				ans = append(ans, node)
			}
		}
	}
	if len(ans) == 0 {
		ans = append(ans, nil)
	}
	return ans
}

func main() {
	ans := generateTrees(3)
	for _, t := range ans {
		fmt.Println(t.Val)
	}
}
