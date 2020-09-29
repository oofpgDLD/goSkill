package array

func moveZeroes(nums []int)  {
	offset := 0
	for i, num := range nums {
		//移动offset个位置
		nums[i-offset] = num
		if num == 0 {
			offset++
		}
	}

	for i := len(nums)-offset; i < len(nums); i++ {
		nums[i] = 0
	}
}