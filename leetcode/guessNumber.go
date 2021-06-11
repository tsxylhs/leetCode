package leetcode

//猜数字游戏的规则如下：
//
//每轮游戏，我都会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。
//如果你猜错了，我会告诉你，你猜测的数字比我选出的数字是大了还是小了。
//你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：
//
//-1：我选出的数字比你猜的数字小 pick < num
//1：我选出的数字比你猜的数字大 pick > num
//0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num
//返回我选出的数字。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/guess-number-higher-or-lower
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func guessNumber(n int) int {
	val := 0
	intlist := make([]int, 1)
	for i := 1; i <= n; i++ {
		intlist = append(intlist, i)
	}
	for {
		g := guess(intlist[len(intlist)/2])
		if g == -1 {
			intlist = intlist[0 : len(intlist)/2]
		} else if g == 1 {
			intlist = intlist[len(intlist)/2 : len(intlist)-1]
		} else if g == 0 {
			val = intlist[len(intlist)/2]
			break
		}
	}
	return val
}

func guessNumber1(n int) int {
	for {
		g := guess(n / 2)
		if g == -1 {
			n = n/2 + ((n - n/2) / 2)
		} else if g == 1 {
			n = n / 2
		} else if g == 0 {

		}
	}
}
func guessNumber2(n int) int {
	left := 1
	right := n
	for {
		num := (left + right) / 2
		switch guess(num) {
		case -1:
			right = num - 1
		case 1:
			left = num + 1
		case 0:
			return num
		}
	}

}
