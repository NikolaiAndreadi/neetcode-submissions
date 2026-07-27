func productExceptSelf(nums []int) []int {
	// 1 2 4 6

	// 2*4*6 1*4*6 1*2*6 1*2*4
	
	//   [1]      1      1*2   1*2*4 
    //  6*4*2   6*4      6      [1]

	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	for i := 0; i < len(nums)-1; i++ {
		res[i+1] = res[i]*nums[i]
	}

	acc := 1
	for i := len(nums)-1; i > 0; i-- {
		res[i] *= acc
		acc *= nums[i]
	}
	res[0] *= acc
	
	return res
}
