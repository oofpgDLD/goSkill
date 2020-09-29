package string

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	temp := make(map[byte]int)
	for _,b := range t{
		/*if v,ok := temp[byte(b)];ok{
			temp[byte(b)] = v+1
		}else {
			temp[byte(b)] = 1
		}*/
		temp[byte(b)] +=1
	}
	for _,b := range s {
		temp[byte(b)] -= 1
		if temp[byte(b)]<0 {
			return false
		}
	}
	return true
}