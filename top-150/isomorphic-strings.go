

func isIsomorphic(s string, t string) bool {
	char_map := make(map[rune]rune)
	trune := []rune(t)
	for index, char := range s {
		val, exists := char_map[char]
		if !exists {
			for _, rch := range char_map {
				if rch == trune[index] {
					return false
				}
			}
			char_map[char] = trune[index]
			continue
		}
		if val == trune[index] {
			continue
		} else {
			return false
		}
	}
	return true
}
