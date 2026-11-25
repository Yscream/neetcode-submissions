func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var result []int
	for i := 0; i < k; i++ {
		maxFreq := 0
		maxKey := 0

		for key, value := range m {
			if value > maxFreq {
				maxFreq = value
				maxKey = key
			}
		}

		result = append(result, maxKey)
		delete(m, maxKey)
	}
	return result
}
