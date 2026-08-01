func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqueList := make(map[int]struct{})
	for _, num := range nums {
		if _, found := uniqueList[num]; !found {
			uniqueList[num] =struct{}{}
		}
	}
	
	list := []int{}
	for num := range uniqueList{
		list = append(list, num)
	}

	sort.Slice(list, func(a, b int) bool {
		return list[a] < list[b]
	})

	streak, maxStreak := 1, 1
	for i:=1 ; i < len(list); i++{
		if list[i] - list[i-1] ==1 {
			streak++
		}else{
			streak=1
		}
		if streak > maxStreak {
            maxStreak = streak
        }
	}
	
	return maxStreak
}
