func isPalindrome(s string) bool {
	input := strings.TrimSpace(strings.ToLower(s))

	input = strings.Map(func(r rune) rune{
		if unicode.IsNumber(r) || unicode.IsLetter(r){
			return r
		}
		return -1
	}, input)
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
