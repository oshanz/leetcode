func strStr(haystack string, needle string) int {

	haystackrunes := []rune(haystack)
	needleLen := len(needle)
	haystackLen := len(haystack)
    needleFirstChar := []rune(needle)[0]
   // fmt.Println(needleFirstChar)
	for i, char := range haystackrunes {
		if haystackLen-i < needleLen {
			return -1
		}
		if char != needleFirstChar {
            continue
        }
		// fmt.Println("i ",i, " l ", needleLen)
		if string(haystackrunes[i:i+needleLen]) == needle {
			return i
		}
		// }
	}
	return -1
}
