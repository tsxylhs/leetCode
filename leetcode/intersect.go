package leetcode

//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//获取两个数组的交集
func Intersect(nums1 []int, nums2 []int) []int {

	ans := make([]int, 0)
	for _, v := range nums1 {
		for k, v2 := range nums2 {
			if v == v2 {
				nums2 = append(nums2[:k], nums2[k+1:]...)
				ans = append(ans, v)
				break
			}
		}
	}

	return ans
}
