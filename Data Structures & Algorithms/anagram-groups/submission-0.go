func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	m := make(map[string][]string)
	for _, s := range strs {
		sortedString := bSort(s)
		m[sortedString] = append(m[sortedString], s)
	}

	out := [][]string{}
	for _, v := range m {
		out = append(out, v)
	}

	return out
}

func bSort(s string) string {
    chars := []rune(s)
    n := len(chars)

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if chars[j] > chars[j+1] {
                chars[j], chars[j+1] = chars[j+1], chars[j]
            }
        }
    }

    return string(chars)
}