package leetcode

import "strconv"

//给你一个 正 整数 num ，输出它的补数。补数是对该数的二进制表示取反。
func FfindComplement(num int) int {
	str := converToBianry(num)
	b := make([]string, 0)
	rev := 0
	for i := 0; i < len(str); i++ {
		v := string(str[i])
		if v == "1" {
			b = append(b, "0")
		} else if v == "0" {
			b = append(b, "1")
		}
	}

	for i, v := range b {
		if v == "1" {
			rev = rev + 1<<(len(b)-i-1)
		}
	}
	return rev
}
func converToBianry(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
