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

	sRune, tRune := []rune(s), []rune(t)

	sort.Slice(sRune, func(i,j int) bool {
		return sRune[i] < sRune[j]
	})

	sort.Slice(tRune, func( i,j int) bool {
		return tRune[i] < tRune[j]
	})	

	for i := range sRune{
		if sRune[i] != tRune[i]  {
			return false
		}
	}

	return true

}
