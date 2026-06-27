func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

	for _, num := range(nums){
		if !seen[num]{
			seen[num] = true
		}else{
			return true
		}
	}
	return false
}
