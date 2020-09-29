package array

func rotate(nums []int, k int) {
	l:=len(nums)
	k %= l
	if k == 0 {
		return
	}
	temps := make([]int, k)
	copy(temps,nums[:k])
	for i := 0; i <= len(nums)-k; i+=k {
		for j := 0; j < k; j++ {
			nums[(i+j+k)%l],temps[j] = temps[j],nums[(i+j+k)%l]
		}
	}
	for n:=0;n<l%k;n++{
		nums[(k-l%k)+n]=temps[n]
	}
}