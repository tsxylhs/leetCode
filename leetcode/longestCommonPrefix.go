package leetcode

import "strconv"

//编写一个函数来查找字符串数组中的最长公共前缀。

func LongestCommonPrefix(strs []string) (str string) {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	testSlice := make([]string, 0)
	vl := make(map[string]int, 1)
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			str := string(strs[i][j]) + strconv.Itoa(j)

			if vl[str] == 0 {
				testSlice = append(testSlice, str)
				vl[str] = 1
			} else {
				vl[str] = vl[str] + 1
			}
		}
	}
	//for _,val:=range strs {
	//     for i,v:=range val{
	//		str:= string(v)+strconv.Itoa(i)
	//     	if vl[str]==0 {
	//			vl[str] = 1
	//		}else{
	//			vl[str]=vl[str]+1
	//		}
	//
	//	 }
	//}
	i := 0
	for j, Key := range testSlice {
		if i != j {
			return str
		}
		if Value, ok := vl[Key]; ok {
			if Value == len(strs) {
				i++
				str = str + string(Key[0])
			}
		}
	}
	return str
}
