package leetcode

import "sort"

//给定一个长度为 n 的 非空 整数数组，每次操作将会使 n - 1 个元素增加 1。找出让数组所有元素相等的最小操作次数
//暴力超时

func MinMoves(nums []int) int {
	n := 1

	sort.Ints(nums)
	if nums[0] == nums[len(nums)-1] {
		return 0
	}
	sorts(nums, &n)
	return n
}

func sorts(nums []int, n *int) {

	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i] + 1
	}
	sort.Ints(nums)
	if nums[0] == nums[len(nums)-1] {
		return
	} else {
		*n = (*n) + 1
		sorts(nums, n)
	}

}
func MinMoves1(nums []int) int {
	sort.Ints(nums)
	n := 0
	for i := len(nums) - 1; i > 0; i-- {
		n = n + nums[i] - nums[0]
	}
	return n
}
