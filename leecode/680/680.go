package _680

func validPalindrome(s string) bool {
	low,high := 0, len(s) -1

	for low < high {
		if s[low] == s[high] {
			low++
			high--
		}else {
			flag := true
			for i, j := low, high-1 ;i<j; i,j = i+1,j-1{
				if s[i] != s[j]{
					flag = false
					break
				}
			}
			if !flag {
				flag = true
				for i, j := low+1, high ;i<j; i,j = i+1,j-1{
					if s[i] != s[j]{
						flag = false
						break
					}
				}
			}
			return flag
		}
	}

	return true
}