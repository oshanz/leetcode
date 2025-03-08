func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if length == 0 {
				continue
			} else {
				break
			}
		}
		length++
	}
	return length
}
