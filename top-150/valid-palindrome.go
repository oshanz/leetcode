func isPalindrome(s string) bool {
	runes := []rune(s)
	j := len(runes) - 1
	i := 0
	for i <= j {
		char_start := runes[i]
		if !(unicode.IsDigit(char_start) || unicode.IsLetter(char_start)) {
			i++
			continue
		}
		char_end := runes[j]
		if !(unicode.IsDigit(char_end) || unicode.IsLetter(char_end)) {
			j--
			continue
		}
		if unicode.ToLower(char_start) != unicode.ToLower(char_end) {
			return false
		}
		i++
		j--
	}
	return true
}
