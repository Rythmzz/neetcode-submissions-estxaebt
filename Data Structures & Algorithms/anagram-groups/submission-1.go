func groupAnagrams(strs []string) [][]string {
	anaGrams := map[string]string{}
	duplicateAnaGrams := map[string]bool{}
	var result [][]string
	for _, str := range strs {
		sortStr := []rune(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})
		anaGrams[str] = string(sortStr)
	}

	for i := 0; i < len(strs); i++ {
		array := []string{strs[i]}
		if duplicateAnaGrams[strs[i]] {
			continue
		} else {
			duplicateAnaGrams[strs[i]] = true
		}
		for j := i + 1; j < len(strs); j++ {
			if anaGrams[strs[i]] == anaGrams[strs[j]] && !duplicateAnaGrams[strs[j]] {
				array = append(array, strs[j])
				duplicateAnaGrams[strs[j]] = true
			} else if strs[i] == strs[j]{
				array = append(array, strs[j])
			}
		}
		result = append(result, array)
	}

	return result
}
