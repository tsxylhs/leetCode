package leetcode

import "sort"

//给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
func ThirdMax(nums []int) int {
	s := RemoveRepeatedElement(nums)
	sort.Ints(s)

	if len(s) > 2 {
		return s[len(s)-3]
	} else {
		return s[len(s)-1]
	}

}

func RemoveRepeatedElement(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
