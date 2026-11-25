func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1, m2 := make(map[uint8]int), make(map[uint8]int)

	for i := range len(s) {
		if _, ok := m1[s[i]]; ok{
			m1[s[i]]++
		} else {
			m1[s[i]] = 1
		}

		if _, ok := m2[t[i]]; ok{
			m2[t[i]]++
		} else {
			m2[t[i]] = 1
		}
	}

	for k, vA := range m1 {
		if vB, ok := m2[k]; !ok || vA != vB {
			return false 
		}
	}
	return true 
}
