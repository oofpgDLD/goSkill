package string

import "math"

func myAtoi(str string) int {
	prefix := ' '
	num := 0
	for _,b := range str{
		if a := b-'0'; a >= 0 && a <= 10 {
			num = num*10+int(a)
			if prefix == ' ' {
				prefix = '+'
			}
			if prefix == '+' && num > math.MaxInt32 {
				return math.MaxInt32
			}
			if prefix == '-' && -num < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			if prefix != ' ' {
				break
			}
			if b == ' ' {
				continue
			}
			if b == '+' || b == '-' {
				prefix = b
			}else {
				return 0
			}
		}
	}
	if prefix == '-' {
		return -num
	}
	return num
}