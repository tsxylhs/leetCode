package leetcode

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func FirstUniqChar(s string) int {
	m := make(map[int32]int, 1)
	for i, v := range s {
		for key, _ := range m {
			if v == key {
				m[key] = i
			}
		}
		m[v] = i
	}
	for i, v := range s {
		if i == m[v] {
			return i
		} else {
			delete(m, v)
		}
	}
	return -1
}
