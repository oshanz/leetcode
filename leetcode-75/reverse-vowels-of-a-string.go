func reverseVowels(s string) string {

	vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	found := []rune{}
	for _, char := range s {
		if slices.Contains(vowels, char) {
			found = append(found, char)
		}
	}
	reversedSlice := []rune{}
	for _, char := range s {
		if slices.Contains(vowels, char) {
			lastIndex := len(found) - 1
			lastChar := found[lastIndex]
			found = found[:lastIndex]
			reversedSlice = append(reversedSlice, lastChar)
		} else {
			reversedSlice = append(reversedSlice, char)
		}
	}
	return string(reversedSlice)
}
