func strStr(haystack string, needle string) int {

	haystackrunes := []rune(haystack)
	needleLen := len(needle)
	haystackLen := len(haystack)
	for i, _ := range haystackrunes {
		if haystackLen-i < needleLen {
			return -1
		}
		// if char == needlerunes[0] {
		// fmt.Println("i ",i, " l ", needleLen)
		if string(haystackrunes[i:i+needleLen]) == needle {
			return i
		}
		// }
	}
	return -1
}
