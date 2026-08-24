func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	input := reg.ReplaceAllString(s, "")	
	
	left, right := 0, len(input)-1
	for left < len(input) && right >0 {
		if input[left] != input[right]{
			return false
		}
		left++
		right--
	}
	return true
}
