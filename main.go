package main

import (
	"fmt"
	"leetcode/leetcode"
)

func main() {
	tree := &leetcode.TreeNode{
		Val: 2,
		Left: &leetcode.TreeNode{
			Val: 0,
			Left: &leetcode.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &leetcode.TreeNode{
			Val: 0,
			Left: &leetcode.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &leetcode.TreeNode{
				Val:   323,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//nums1 := "aadfaf"
	//i := []int{1, 2, 3}
	//i := []int{1, 2, 2, 5, 3, 5}
	//fmt.Println(leetcode.FirstUniqChar("lleetcode"))
	//nums2 := "badfafsaes"
	//fmt.Println(leetcode.MinMoves1(i))
	//t := []string{"Hello", "Alaska", "Dad", "Peace"}
	//fmt.Println(leetcode.FindWords(t))
	//fmt.Println(leetcode.FfindComplement(50))
	//fmt.Println(leetcode.ThirdMax(i))
	fmt.Println(leetcode.SumOfLeftLeaves(tree))
	fmt.Println(leetcode.CountBits(10))
}
