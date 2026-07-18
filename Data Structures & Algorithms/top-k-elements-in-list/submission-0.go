func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}
	
	uniqueNums := []int{}
	for num := range freq {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(a,b int) bool {
		return freq[uniqueNums[a]] > freq[uniqueNums[b]]
	})

	final := []int{}
	for i:=0; i < k ; i++{
		final = append(final, uniqueNums[i])
	}
	return final
}