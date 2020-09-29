package array

func containsDuplicate(nums []int) bool {
	temp := make(map[int]bool)
	for _, num := range nums {
		if ok := temp[num];ok {
			return true
		}
		temp[num] = true
	}
	return false
}