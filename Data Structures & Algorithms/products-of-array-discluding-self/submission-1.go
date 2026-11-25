func productExceptSelf(nums []int) []int {
	res1 := make([]int, len(nums))
	resElementProduct1 := 1

	for i := range nums {
		res1[i] = resElementProduct1
		resElementProduct1 *= nums[i]
	}

	res2 := make([]int, len(nums))
	resElementProduct2 := 1
	for i := len(nums)-1; i >= 0; i-- {
		res2[i] = resElementProduct2
		resElementProduct2 *= nums[i]
	}

	for i := range res1 {
    	res1[i] *= res2[i]
	}

	return res1
}
