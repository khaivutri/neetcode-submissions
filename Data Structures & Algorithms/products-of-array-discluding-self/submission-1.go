func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums), len(nums))
	startPrefix := 1
	for i, num := range nums {
		startPrefix = startPrefix * num
		prefix[i] = startPrefix 
	}

	
	postfix := make([]int,len(nums),len(nums))
	startPostfix := 1
	for i := len(nums) -1; i >=0; i--{
		startPostfix = startPostfix * nums[i]
		postfix[i] = startPostfix
	}

	result := make([]int,len(nums),len(nums))
	for index := range nums{
		if index -1 == -1{
			result[index] = postfix[index+1]
			continue
		}
		if index +1 == len(nums){
			result[index] = prefix[index-1]
			break
		}
		result[index] = prefix[index-1] * postfix[index+1]
	}
	return result
}


