package string

func reverse(x int) int {
	ans := 0
	for x != 0{
		temp := x%10
		ans = ans*10 + temp
		x /= 10
	}
	if ans > 1<<31 - 1 || ans < -1 << 31{
		return 0
	}
	return ans
}