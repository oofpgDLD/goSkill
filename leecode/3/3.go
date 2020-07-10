package _03

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	child := make([]byte,0)
	i, j := 0, 0
	max := 0
	for ;i<len(bs);i++{
		//check exists
		for k:=j;k < len(child);k++{
			if child[k] == bs[i]{
				j = k + 1
				break
			}
		}
		child = append(child, bs[i])
		if len(child[j:]) > max {
			max = len(child[j:])
		}
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	bs := []byte(s)
	child := make(map[byte]int)
	i, j := 0, 0
	max := 0
	for ;i<len(bs);i++{
		//check exists
		if idx,ok := child[bs[i]];ok{
			if idx >= j{
				child[bs[i]] = i
				j = idx + 1
				continue
			}
		}
		child[bs[i]] = i
		if ln := i-j + 1; ln > max {
			max = ln
		}
	}
	return max
}