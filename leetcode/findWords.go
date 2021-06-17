package leetcode

import "strings"

/*
给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。

美式键盘 中：

第一行由字符 "qwertyuiop" 组成。
第二行由字符 "asdfghjkl" 组成。
第三行由字符 "zxcvbnm" 组成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/keyboard-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func FindWords(words []string) []string {
	fm := make(map[int32]int, 0)
	fs := ""
	tm := make(map[int32]int, 0)
	ts := ""
	sm := make(map[int32]int, 0)
	ss := ""
	f := "qwertyuiop"
	s := "asdfghjkl"
	t := "zxcvbnm"
	reVal := make([]string, 0)
	for _, v := range f {
		fm[v] = 1
	}
	for _, v := range s {
		sm[v] = 1
	}
	for _, v := range t {
		tm[v] = 1
	}
	for _, word := range words {
		wordt := strings.ToLower(word)
		for _, val := range wordt {

			if fm[val] == 1 {
				fs = fs + string(val)
			} else if sm[val] == 1 {
				ss = ss + string(val)
			} else if tm[val] == 1 {
				ts = ts + string(val)
			}
		}
		if wordt == fs || wordt == ss || wordt == ts {
			reVal = append(reVal, word)
		}

		ts, ss, fs = "", "", ""
	}
	return reVal
}
