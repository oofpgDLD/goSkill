package array

func singleNumber(nums []int) int {
	temp := make(map[int]int)
	for _, num := range nums {
		temp[num]++
	}
	for k,v := range temp{
		if v < 2 {
			return k
		}
	}
	return 0
}