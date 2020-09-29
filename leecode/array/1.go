package array

func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums);i++{
		currency := nums[i]
		cnt := 0
		for j := i+1; j < len(nums)-count;j++{
			nums[j-cnt] = nums[j]
			if nums[j] == currency {
				cnt++
			}
		}
		count += cnt
	}
	nums = nums[:len(nums)-count]
	return len(nums)
}