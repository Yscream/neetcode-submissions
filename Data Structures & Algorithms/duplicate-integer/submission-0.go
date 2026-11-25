
func hasDuplicate(nums []int) bool {
    for i := range nums {
        s := nums[i+1:]
        for j := range s {
            if nums[i] == s[j] {
                return true
            }
        }
    }
    return false
}
