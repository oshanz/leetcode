func canConstruct(ransomNote string, magazine string) bool {
	ransomLetterCountMap := letterCount(ransomNote)
    magazineLetterCountMap := letterCount(magazine)

    for ransomLetter, count := range ransomLetterCountMap{
        if count > magazineLetterCountMap[ransomLetter]{
            return false
        }
    }
    return true
}

func letterCount(letters string) map[rune]int {
	letterCountMap := make(map[rune]int)
	for _, letter := range letters {
		letterCountMap[letter] += 1
	}
	return letterCountMap
}
