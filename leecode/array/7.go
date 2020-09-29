package array

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		temp := digits[i]+carry
		digits[i] = temp%10
		carry = temp/10
	}
	if carry == 1 {
		ret := make([]int,1,len(digits)+1)
		ret = append(ret, digits...)
		ret[0] = 1
		return ret
	}
	return digits
}