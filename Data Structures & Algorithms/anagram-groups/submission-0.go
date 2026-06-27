

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 1 || len(strs) >1000{
		return nil
	}

	anagram := make(map[string][]string)

	for _, str := range strs{
		strRune := []rune(str)

		sort.Slice(strRune, func ( i, j int) bool {
			return strRune[i] < strRune[j]
		})
		sortedStr := string(strRune)
		
		if _, found := anagram[sortedStr]; !found{
			anagram[sortedStr] = []string{}
		}

		anagram[sortedStr] = append(anagram[sortedStr], str)

	}

	var output [][]string

	for _, value := range anagram{
		output = append(output, value)
	}
	return output
}
