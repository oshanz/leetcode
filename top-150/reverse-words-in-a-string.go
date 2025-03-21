func reverseWords(s string) string {

	wordEnd := -1
	// reversedWord := ""
	var reversedWord strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if wordEnd == -1 {
			if s[i] != ' ' { // start of a word
				wordEnd = i
			}
		}
		if wordEnd != -1 {
			if i == 0 { //end of the string
				// reversedWord += s[i : wordEnd+1]
				reversedWord.WriteString(s[i : wordEnd+1])
				wordEnd = -2
			} else if s[i-1] == ' ' { //end of a word
				// reversedWord += s[i:wordEnd+1] + " "
				reversedWord.WriteString(s[i : wordEnd+1])
				reversedWord.WriteString(" ")
				wordEnd = -1
			}
		}
	}
	return strings.TrimSpace(reversedWord.String())
	// if wordEnd == -1 {
	// 	return reversedWord[:len(reversedWord)-1] // remove last space
	// }
	// return reversedWord
}
