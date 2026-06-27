
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	if len(s) < 1 || len(s) > 5*10e4{
		return false
	}
	if len(t) < 1 || len(t) > 5*10e4{
		return false
	}

	freq := make(map[rune]int)
	for _, char := range s{
		freq[char]++
	}

	for _, char := range t{
		if count, found := freq[char]; !found || count ==0{
			return false
		}
		freq[char]--
	
	}
	
	// for _, count := range freq{
	// 	if count >0 {
	// 		return false
	// 	}
	// }
	return true
}
