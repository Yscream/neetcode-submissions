func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)

	for _, v := range s {
		if _, ok := m1[v]; ok{
			m1[v]++
			continue
		}
		m1[v] = 1
	}

	m2 := make(map[rune]int)
	for _, v := range t {
		if _, ok := m2[v]; ok{
			m2[v]++
			continue
		}
		m2[v] = 1
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, vA := range m1 {
		if vB, ok := m2[k]; !ok || vA != vB {
			return false 
		}
	}
	return true 
}
