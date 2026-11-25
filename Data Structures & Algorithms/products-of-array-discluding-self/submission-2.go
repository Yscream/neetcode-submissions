func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	leftProduct := 1

	for i := range nums {
		res[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := len(nums)-1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
